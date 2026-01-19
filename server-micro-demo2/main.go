package main

import (
	"greeter/handler"
	pb "greeter/proto"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/micro/plugins/v5/registry/consul"
	"go-micro.dev/v5/server"

	"go-micro.dev/v5"
)

var (
	name    = "greeter"
	version = "latest"
)

func main() {
	consulCfg := consulapi.DefaultConfig()
	consulCfg.Address = "127.0.0.1:8500"
	registry := consul.NewRegistry(consul.Config(consulCfg))
	// Create service
	service := micro.NewService(
		micro.Name(name),
		micro.Version(version),
		micro.Server(server.NewServer(
			server.Name(name),
			server.Registry(registry),
			server.Address("0.0.0.0:9099"),
			server.Advertise("127.0.0.1:9099"),
		)),
	)
	// Initialize service
	service.Init()

	// Register handler
	pb.RegisterGreeterHandler(service.Server(), handler.New())

	// Run service
	service.Run()
}
