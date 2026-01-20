package it

import (
	"gin-demo/models"
	pb "gin-demo/proto"
	"github.com/gin-gonic/gin"
	log "go-micro.dev/v5/logger"

	"net/http"
)

type GoodsController struct{}

func (con GoodsController) Index(c *gin.Context) {
	goodsService := pb.NewGoodsService("goods", models.MicroClient)

	goods, err := goodsService.AddGoods(c.Request.Context(), &pb.AddGoodsReq{
		GoodsModel: &pb.GoodsModel{
			Title:   "添加商品",
			Price:   1000,
			Content: "这是一个商品",
		},
	})
	if err != nil {
		log.Fatalf("add goods err: %v", err)
	}
	log.Info(goods)

	c.JSON(http.StatusOK, gin.H{
		"message": goods.Msg,
		"code":    goods.Code,
	})

}
func (con GoodsController) GetGoods(c *gin.Context) {
	goodsService := pb.NewGoodsService("goods", models.MicroClient)
	goods, err := goodsService.GetGoods(c.Request.Context(), &pb.GetGoodsReq{})
	if err != nil {
		log.Fatalf("get goods err: %v", err)
	}
	log.Info(goods)
	c.JSON(http.StatusOK, gin.H{
		"data": goods,
	})
}
