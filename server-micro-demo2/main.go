package main

import (
	"greeter/handler"
	pb "greeter/proto"

	"github.com/micro/plugins/v5/registry/consul"

	"go-micro.dev/v5"
)

var (
	name    = "greeter"
	version = "latest"
)

func main() {
	registry := consul.NewRegistry()
	// Create service
	service := micro.NewService(
		micro.Name(name),
		micro.Version(version),
		micro.Registry(registry))
	// Initialize service
	service.Init()

	// Register handler
	pb.RegisterGreeterHandler(service.Server(), handler.New())

	// Run service
	service.Run()
}
