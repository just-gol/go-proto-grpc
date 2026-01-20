package handler

import (
	"context"
	"strconv"

	log "go-micro.dev/v5/logger"

	pb "goods/proto"
)

type Goods struct{}

func New() *Goods {
	return &Goods{}
}

func (e *Goods) AddGoods(ctx context.Context, req *pb.AddGoodsReq, resp *pb.AddGoodsResp) error {
	log.Info("Received Goods.Call request", req)
	resp.Msg = "添加成功"
	resp.Code = "200"
	log.Info("Received Goods.Call response", resp)
	return nil
}
func (e *Goods) GetGoods(ctx context.Context, req *pb.GetGoodsReq, resp *pb.GetGoodsResp) error {
	var list []*pb.GoodsModel
	for i := 0; i < 10; i++ {
		list = append(list, &pb.GoodsModel{
			Title:   "商品" + strconv.Itoa(i),
			Price:   float64(i),
			Content: "好商品" + strconv.Itoa(i),
		})
	}
	resp.GoodsList = list
	log.Info("Received Goods.Call request", req)
	return nil
}
