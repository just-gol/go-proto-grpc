package main

import (
	"github.com/micro/plugins/v5/registry/consul"
	"service/handler"
	pb "service/proto"

	"go-micro.dev/v5"
)

var (
	service = "service-micro-demo1"
	version = "latest"
)

func main() {
	//集成consul
	registry := consul.NewRegistry()
	// Create service-micro-demo1
	service := micro.NewService(
		// 注册consul
		micro.Registry(registry),
		micro.Version(version),
		micro.Name(service),
	)

	// Initialize service-micro-demo1
	service.Init()

	// Register handler
	pb.RegisterServiceHandler(service.Server(), handler.New())

	// Run service-micro-demo1
	service.Run()
}
