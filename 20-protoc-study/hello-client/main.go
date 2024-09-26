package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	proto "protoc-test/hello-client/proto"
)

type ClienTokenAuth struct {
}

func (c ClienTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		// 这里key的字母大小写传入到服务端都会变成小写
		// map[:authority:[127.0.0.1:9090] appid:[ops] appkey:[123456]
		"appid":  "ops",
		"appkey": "123456",
	}, nil
}

func (c ClienTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	//creds, _ := credentials.NewClientTLSFromFile("D:\\0_boke_work\\2-vscode-dir\\1-local-test-dir\\protoc-test\\key\\test.pem", "*.ishells.cn")

	// 连接到server端，此处禁用安全传输，没有加密和验证(后面加上了credentials的token认证)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClienTokenAuth)))

	conn, err := grpc.NewClient("127.0.0.1:9090", opts...)
	//conn, err := grpc.NewClient("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//加密认证未生效
	//conn, err := grpc.NewClient("127.0.0.1:9090", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 建立连接
	client := proto.NewSayHelloClient(conn)
	// 执行rpc调用（这个方法在服务器端来实现并返回结果）
	resp, _ := client.SayHello(context.Background(), &proto.HelloRequest{
		RequestName: "ops",
	})
	fmt.Println(resp.GetResponseMsg())
}
