package main

import (
	"context"
	"fmt"
	"greeter/proto/greeter"
	"net"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

type Hello struct {
	greeter.UnimplementedGreetServiceServer
}

func (h Hello) SayHello(c context.Context, req *greeter.HelloReq) (*greeter.HelloResp, error) {
	fmt.Println("call Hello.SayHello:", req)
	return &greeter.HelloResp{
		Message: "Hello " + req.Name,
	}, nil
}

func main() {
	// 注册consul服务
	//初始化consul配置
	config := api.DefaultConfig()
	// 获取consul操作对象
	client, err := api.NewClient(config)
	if err != nil {
	}
	// 3、配置注册服务的参数
	agentService := api.AgentServiceRegistration{
		ID:      "1",
		Tags:    []string{"test"},
		Name:    "HelloService",
		Port:    8080,
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			TCP:      "127.0.0.1:8080",
			Timeout:  "5s",
			Interval: "30s",
		},
	}
	//注册服务到consul上
	client.Agent().ServiceRegister(&agentService)

	server := grpc.NewServer()
	greeter.RegisterGreetServiceServer(server, new(Hello))
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
	}
	defer listen.Close()
	err = server.Serve(listen)
	if err != nil {
	}

}
