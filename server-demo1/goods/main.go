package main

import (
	"context"
	"fmt"
	"goods/proto/goods"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type Goods struct {
	goods.UnimplementedGoodsServer
}

func (g Goods) AddGoods(c context.Context, in *goods.AddGoodsReq) (*goods.AddGoodsRes, error) {
	fmt.Println(in)
	return &goods.AddGoodsRes{
		Message: "success",
		Success: true,
	}, nil
}
func (g Goods) GetGoods(c context.Context, in *goods.GetGoodsReq) (*goods.GetGoodsRes, error) {
	fmt.Println(in)
	var list []*goods.GoodsModel
	for i := 0; i < 10; i++ {
		list = append(list, &goods.GoodsModel{
			Title:   "商品" + strconv.Itoa(i),
			Price:   float64(i),
			Content: "测试商品内容" + strconv.Itoa(i),
		})
	}
	return &goods.GetGoodsRes{
		GoodsList: list,
	}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	goods.RegisterGoodsServer(grpcServer, &Goods{})
	listen, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {

	}
	_ = grpcServer.Serve(listen)
}
