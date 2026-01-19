package main

import (
	"client-demo2/proto/goods"
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
		panic("no server-micro-demo1 entry for HelloService with tag test")
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
		panic(err)
	}
	fmt.Printf("%v \n", resp.Message)

	//获取goods服务
	goodsEntry, _, _ := client.Health().Service("GoodsService", "goods", false, nil)
	goodsAddress := goodsEntry[0].Service.Address + ":" + strconv.Itoa(goodsEntry[0].Service.Port)
	goodsConn, err := grpc.NewClient(goodsAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	goodsClient := goods.NewGoodsClient(goodsConn)
	addGoods, err := goodsClient.AddGoods(context.Background(), &goods.AddGoodsReq{
		Goods: &goods.GoodsModel{
			Title:   "测试商品",
			Price:   20,
			Content: "测试商品的内容"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", addGoods.Message)
	getGoods, err := goodsClient.GetGoods(context.Background(), &goods.GetGoodsReq{})
	if err != nil {
	}
	fmt.Printf("%v", getGoods)

}
