package main

import (
	"client-demo2/proto/greeter"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"strconv"
)

func main() {
	// 初始化consul配置
	config := api.DefaultConfig()
	// 获取consul操作对象
	client, err := api.NewClient(config)
	if err != nil {
	}
	// 获取地址
	serviceEntry, _, _ := client.Health().Service("HelloService", "test", false, nil)
	fmt.Println(serviceEntry[0].Service.Address)
	fmt.Println(serviceEntry[0].Service.Port)
	address := serviceEntry[0].Service.Address + ":" + strconv.Itoa(serviceEntry[0].Service.Port)

	conn, err := grpc.NewClient(address)
	if err != nil {
	}
	serviceClient := greeter.NewGreetServiceClient(conn)
	resp, err := serviceClient.SayHello(context.Background(), &greeter.HelloReq{
		Name: "HelloService",
	})
	if err != nil {
	}
	fmt.Println(resp)

}
