package main

import (
	"context"
	"fmt"
	pb "goods/proto"

	"github.com/micro/plugins/v5/registry/consul"
	"go-micro.dev/v5"
)

func main() {
	registry := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("goods-client"),
		micro.Registry(registry),
	)

	// Initialize service
	service.Init()

	// Register handler
	goodsService := pb.NewGoodsService("goods", service.Client())
	goods, err := goodsService.AddGoods(context.Background(), &pb.AddGoodsReq{
		Title:   "测试",
		Price:   1000,
		Content: "测试12",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(goods)

}
