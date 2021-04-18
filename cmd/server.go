package cmd

import (
	"context"
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	rd "github.com/go-redis/redis/v8"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/viper"
	"github.com/vegggg/user-management/handler"
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
	return r
}

func newMysqlConn(address string) (*sql.DB, error) {
	return sql.Open("mysql", address)
}

func startHttpGateway() error {
	mux := runtime.NewServeMux()
	ctx := context.Background()

	// dependencies injection starts here
	mysqlDB, err := newMysqlConn(viper.GetString("mysql.conn_address"))
	if err != nil {
		panic("connection to db failed:" + err.Error())
	}

	mr := sqlrepo.NewMysqlRepo(mysqlDB)

	user_managementpb.
		RegisterUserManagementHandlerServer(
			ctx,
			mux,
			handler.NewUserManagement(mr))

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
	return nil
}
