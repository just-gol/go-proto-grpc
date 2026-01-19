package main

import (
	"context"
	"fmt"
	pb "greeter/proto"

	"github.com/micro/plugins/v5/registry/consul"
	"go-micro.dev/v5"
	log "go-micro.dev/v5/logger"
)

var (
	name    = "greeter"
	version = "latest"
)

func main() {
	// 集成consul
	registry := consul.NewRegistry()
	// Create service
	service := micro.NewService(
		micro.Name(name),
		micro.Version(version),
		micro.Registry(registry),
	)

	// Initialize service
	service.Init()

	// Register handler
	greeterService := pb.NewGreeterService(name, service.Client())
	call, err := greeterService.Call(context.Background(), &pb.Request{
		Name: "hello",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(call)

	// Run service
	service.Run()
}
