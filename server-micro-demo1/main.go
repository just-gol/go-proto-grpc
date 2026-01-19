package main

import (
	"service/handler"
	pb "service/proto"

	"github.com/micro/plugins/v5/registry/consul"

	"go-micro.dev/v5"
)

var (
	service = "server-micro-demo1"
	version = "latest"
)

func main() {
	//集成consul
	registry := consul.NewRegistry()
	// Create server-micro-demo1
	service := micro.NewService(
		// 注册consul
		micro.Registry(registry),
		micro.Version(version),
		micro.Name(service),
	)

	// Initialize server-micro-demo1
	service.Init()

	// Register handler
	pb.RegisterServiceHandler(service.Server(), handler.New())

	// Run server-micro-demo1
	service.Run()
}
