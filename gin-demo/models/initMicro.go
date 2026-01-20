package models

import (
	"github.com/micro/plugins/v5/registry/consul"
	"go-micro.dev/v5"
	"go-micro.dev/v5/client"
)

var MicroClient client.Client

func init() {
	registry := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("goods-client"),
		micro.Registry(registry),
	)

	// Initialize service
	service.Init()

	MicroClient = service.Client()
}
