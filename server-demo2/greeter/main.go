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
	config.Address = "http://127.0.0.1:8500"
	// 获取consul操作对象
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	serviceIP := "127.0.0.1"
	servicePort := 9090
	// 3、配置注册服务的参数
	agentService := api.AgentServiceRegistration{
		ID:      "1",
		Tags:    []string{"test"},
		Name:    "HelloService",
		Port:    servicePort,
		Address: serviceIP,
		Check: &api.AgentServiceCheck{
			TCP:      serviceIP + ":9090",
			Timeout:  "5s",
			Interval: "30s",
		},
	}
	//注册服务到consul上
	if err := client.Agent().ServiceRegister(&agentService); err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	greeter.RegisterGreetServiceServer(server, new(Hello))
	listen, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}

}
