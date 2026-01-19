package main

import (
	"goods/handler"
	pb "goods/proto"

	"github.com/micro/plugins/v5/registry/consul"
	"go-micro.dev/v5"
	"go-micro.dev/v5/server"
)

var (
	name    = "goods"
	version = "0.0.1"
)

func main() {
	registry := consul.NewRegistry()
	// Create service
	service := micro.NewService(
		micro.Name(name),
		micro.Registry(registry),
		micro.Version(version),
		micro.Server(server.NewServer(
			server.Name(name),
			server.Registry(registry),
			server.Address("0.0.0.0:9098"),
			server.Advertise("127.0.0.1:9098"),
		)),
	)

	// Initialize service
	service.Init()

	// Register handler
	pb.RegisterGoodsHandler(service.Server(), handler.New())

	// Run service
	service.Run()
}
