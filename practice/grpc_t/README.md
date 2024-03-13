# 准备环境
## 下载
protoc 安装地址：[下载](https://objects.githubusercontent.com/github-production-release-asset-2e65be/23357588/71a52f62-f8be-4e9b-a8c8-dcc0da0e582d?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAVCODYLSA53PQK4ZA%2F20240125%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240125T144027Z&X-Amz-Expires=300&X-Amz-Signature=f64d47813d10aa45e15d7c4465e45d5974e3e10e22e2228c5f5f0ab9074e6e46&X-Amz-SignedHeaders=host&actor_id=104589189&key_id=0&repo_id=23357588&response-content-disposition=attachment%3B%20filename%3Dprotoc-25.2-linux-x86_64.zip&response-content-type=application%2Foctet-stream)

## 安装
linux 安装方式：

移动protoc文件到`/usr/local/bin/`目录下

移动include 文件到`/usr/local/`下

## 验证
```shell
protoc --versioin
```

# 文件编写
xxx.protoc 
```protobuf 
// 版本号
syntax = "proto3";

// 指定文件生成的路径。当前路径下包名为service
option go_package = ".;service";

// 定义结构体
message UserRequest {
  // 定义用户名
  string name = 1;
}

// 响应结构体
message UserResponse {
  int32           id    = 1;
  string          name  = 2;
  int32           age   = 3;
  // repeated修饰符是可变数组，go转切片
  repeated string hobby = 4;
}

// service定义方法
service UserInfoService {
  rpc GetUserInfo (UserRequest) returns (UserResponse) {
  }
}
```

# 生成文件
```shell
# 生成go文件
protoc --go_out=. xxxx.proto

# 生成grpc相关文件
protoc --go-grpc_out=. xxx.proto
```