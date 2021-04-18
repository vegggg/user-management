package cmd

import (
	"context"
	"net/http"

	rd "github.com/go-redis/redis/v8"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/viper"
	"github.com/vegggg/user-management/handler"
	user_managementpb "github.com/vegggg/user-management/proto/user_mgnt/v1"
	"google.golang.org/grpc"
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

func startGrpcServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s := grpc.NewServer()

	user_managementpb.RegisterUserManagementServer(s, handler.NewUserManagement())

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	// mux := runtime.NewServeMux()
	// opts := []grpc.DialOption{grpc.WithInsecure()}
	// err := user_managementpb.Re(ctx, mux, s, opts)
	// if err != nil {
	// 	return err
	// }

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	// return http.ListenAndServe(":8081", mux)

	return nil
}

func startHttpGateway() error {
	mux := runtime.NewServeMux()
	ctx := context.Background()

	user_managementpb.RegisterUserManagementHandlerServer(ctx, mux, handler.NewUserManagementAPI())

	s := &http.Server{
		Addr:    ":8080",
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
