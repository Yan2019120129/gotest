package main

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gotest/common/module/logs"
	service "gotest/practice/grpc_test/proto"
	"net"
)

const LogName = "grpcServer"

// UserInfoService 服务实例
type UserInfoService struct {
	service.UnimplementedUserInfoServiceServer
}

// GetUserInfo 实现grpc user 接口
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *service.UserRequest) (*service.UserResponse, error) {
	name := req.Name
	logs.Logger.Info(LogName, zap.String("name", name))

	if name == "yan" {
		return &service.UserResponse{
			Id:    1,
			Name:  "wang",
			Age:   18,
			Hobby: []string{"lan qui", "ping pang"},
		}, nil
	}

	return nil, nil
}

func (s *UserInfoService) mustEmbedUnimplementedUserInfoServiceServer() {}

func main() {
	// 服务地址
	addr := "127.0.0.1:8080"

	// 1. 监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logs.Logger.Error(LogName, zap.Error(err))
	}
	logs.Logger.Info(LogName, zap.String("addr", addr))

	// 2. 实例化grpc
	s := grpc.NewServer()

	// 3.在grpc上注册微服务
	service.RegisterUserInfoServiceServer(s, &UserInfoService{})

	if err = s.Serve(listener); err != nil {
		logs.Logger.Error(LogName, zap.Error(err))
		return
	}
}
