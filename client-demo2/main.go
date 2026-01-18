package main

import (
	"client-demo2/proto/greeter"
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 初始化consul配置
	config := api.DefaultConfig()
	// 获取consul操作对象
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	// 获取地址
	serviceEntry, _, err := client.Health().Service("HelloService", "test", false, nil)
	if err != nil {
		panic(err)
	}
	if len(serviceEntry) == 0 {
		panic("no service entry for HelloService with tag test")
	}
	fmt.Println(serviceEntry[0].Service.Address)
	fmt.Println(serviceEntry[0].Service.Port)
	address := serviceEntry[0].Service.Address + ":" + strconv.Itoa(serviceEntry[0].Service.Port)

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	serviceClient := greeter.NewGreetServiceClient(conn)
	resp, err := serviceClient.SayHello(context.Background(), &greeter.HelloReq{
		Name: "HelloService",
	})
	if err != nil {
	}
	fmt.Println(resp)

}
