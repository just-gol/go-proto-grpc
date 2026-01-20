package handler

import (
	"context"

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
