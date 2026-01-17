package main

import (
	"client-demo1/goods/goods"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
	}
	client := goods.NewGoodsClient(conn)
	//hello, err := client.SayHello(context.Background(), &greeter.HelloReq{
	//	Name: "hello",
	//})
	//if err != nil {
	//}
	//fmt.Printf("%+v\n", hello)
	model := &goods.GoodsModel{
		Title:   "1",
		Price:   10000.0,
		Content: "1",
	}
	res, err := client.AddGoods(context.Background(), &goods.AddGoodsReq{
		Goods: model,
	})
	if err != nil {

	}
	fmt.Printf("%+v\n", res)
	getGoods, err := client.GetGoods(context.Background(), new(goods.GetGoodsReq))
	if err != nil {
	}
	fmt.Printf("%+v\n", getGoods)
}
