module github.com/vegggg/user-management

go 1.15

replace (
	github.com/vegggg/user-management/proto-gen/user_management => /Users/macos/Documents/go-src/user-management
	google.golang.org/grpc => google.golang.org/grpc v1.29.0
)

require (
	github.com/go-redis/redis/v8 v8.8.2
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	go.uber.org/zap v1.10.0
	google.golang.org/genproto v0.0.0-20210416161957-9910b6c460de
	google.golang.org/grpc v1.36.1
	google.golang.org/protobuf v1.26.0
)
