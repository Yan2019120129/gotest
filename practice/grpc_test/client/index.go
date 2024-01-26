package main

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gotest/common/module/logs"
	service "gotest/practice/grpc_test/proto"
)

const LogName = "grpcClient"

// 1.连接服务端
// 2.实例gRPC客户端
// 3.调用

func main() {
	// 1.连接
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		logs.Logger.Info(LogName, zap.Error(err))
	}
	defer conn.Close()

	// 2. 实例化gRPC客户端
	client := service.NewUserInfoServiceClient(conn)

	// 3.组装请求参数
	req := &service.UserRequest{Name: "yan"}

	// 4. 调用接口
	response, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		logs.Logger.Error(LogName, zap.Error(err))
	}
	logs.Logger.Info(LogName, zap.Reflect("相应结果：", response))
}
