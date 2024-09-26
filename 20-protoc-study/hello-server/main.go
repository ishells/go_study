package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"net"
	proto "protoc-test/hello-server/proto"
)

// hello.proto 中的 Service SayHello 方法对应 hello_grpc.pb.go 中定义的 SayHello 方法(并由UnimplementedSayHelloServer结构体实现)
// 所以当我们要在main函数中实现该方法时，需要继承UnimplementedSayHelloServer结构体，并实现SayHello方法
type server struct {
	proto.UnimplementedSayHelloServer
}

// 将hello_grpc.pb.go的SayHello方法拿过来，实现者改为server结构体，参数和返回值改为符合 hello_grpc.pb.go 中 SayHello方法定义的
func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	// 获取客户端传入的元数据
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("未传输token")
	}

	fmt.Println("md: ", md)
	var appid, appkey string
	if v, ok := md["appid"]; ok {
		appid = v[0]
		fmt.Println("md.appid:", md["appid"])
	}
	if v, ok := md["appkey"]; ok {
		appkey = v[0]
		fmt.Println("md.appkey:", md["appkey"])
	}
	if appid != "ops" || appkey != "123456" {
		return nil, errors.New("token验证失败")
	}

	fmt.Println("hello " + req.RequestName)
	return &proto.HelloResponse{ResponseMsg: "hello " + req.RequestName}, nil
}

// 基于 hello_grpc.pb.go 的 SayHello 方法实现之后，就需要暴露一个端口出去提供服务
func main() {
	//creds, _ := credentials.NewServerTLSFromFile("D:\\0_boke_work\\2-vscode-dir\\1-local-test-dir\\protoc-test\\key\\test.pem", "D:\\0_boke_work\\2-vscode-dir\\1-local-test-dir\\protoc-test\\key\\test.key")

	// 开启端口（此处选择忽略报错）
	listen, _ := net.Listen("tcp", ":9090")
	// 创建grpc服务
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	//grpcServer := grpc.NewServer(grpc.Creds(creds))
	// 在grpc服务端中注册我们自己编写的服务
	proto.RegisterSayHelloServer(grpcServer, &server{})

	// 启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("faile to serve: %v", err)
		return
	}
}
