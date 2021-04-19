package cmd

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	rd "github.com/go-redis/redis/v8"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/viper"
	"github.com/vegggg/user-management/handler"
	"github.com/vegggg/user-management/otp"
	user_managementpb "github.com/vegggg/user-management/proto/user_mgnt/v1"
	"github.com/vegggg/user-management/user/sqlrepo"
)

func newRedisClient() *rd.Client {
	addr := viper.GetString("redis.address")
	pass := viper.GetString("redis.password")
	db := viper.GetInt("redis.db")

	r := rd.NewClient(&rd.Options{
		Addr:     addr,
		Password: pass, // no password set
		DB:       db,   // use default DB
	})
	if _, err := r.Ping(context.Background()).Result(); err != nil {
		panic("ping redis timed out")
	}
	return r
}

func newMysqlConn() (*sql.DB, error) {
	address := viper.GetString("mysql.conn_address")
	return sql.Open("mysql", address)
}

func startHttpGateway() error {
	mux := runtime.NewServeMux()
	ctx := context.Background()

	// dependencies injection starts here

	// mysql
	mysqlDB, err := newMysqlConn()
	if err != nil {
		panic("connection to db failed:" + err.Error())
	}
	// redis
	redisConn := newRedisClient()

	user := sqlrepo.NewMysqlRepo(mysqlDB)
	otp := otp.NewRedisOTP(
		redisConn,
		5,
		1*time.Minute,
	)

	user_managementpb.
		RegisterUserManagementHandlerServer(
			ctx,
			mux,
			handler.NewUserManagement(user, otp))

	s := &http.Server{
		Addr:    viper.GetString("service.server_address"),
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		glog.Infof("Shutting down the http gateway server")
		if err := s.Shutdown(context.Background()); err != nil {
			glog.Errorf("Failed to shutdown http gateway server: %v", err)
		}
	}()

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		glog.Errorf("Failed to listen and serve: %v", err)
		return err
	}
	glog.Infof("Started Server")
	return nil
}
