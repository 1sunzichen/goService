package shop_cart

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopro/gin_test/mxshop-api/order-web/api"
	"gopro/gin_test/mxshop-api/order-web/global"
	"gopro/gin_test/mxshop-api/order-web/proto"
	"net/http"
	"strconv"
)

func List(c *gin.Context){
	//获取购物车商品
	userId,_:=strconv.Atoi(c.Query("userId"))

	fmt.Println(userId,"userId",c.Query("userId"))
	rsp,err:=global.OrderClient.CartItemList(context.Background(),&proto.UserInfo{

		Id:int32(userId),
	})
	if err!=nil{
		zap.S().Errorw("[List]查询 【购物车列表】失败")
		api.HandleGrpcErrorToHttp(err,c)
		return
	}
	ids:=make([]int32,0)
	for _,item:=range rsp.Data{
		ids=append(ids,item.GoodsId)
	}
	if len(ids)==0{
		c.JSON(http.StatusOK,gin.H{
			"total":0,
		})
		return
	}
	//请求商品服务获得商品信息
	goodsRsp,err:=global.GoodsClient.BatchGetGoods(context.Background(),&proto.BatchGoodsIdInfo{Id: ids})
	if err!=nil{
		zap.S().Errorw("【list】 查询【 购物车列表失败】")
		api.HandleGrpcErrorToHttp(err,c)
		return
	}
	goodsList:=make([]interface{},0)
	for _,item:=range rsp.Data{
		for _,good:=range goodsRsp.Data{
			if good.Id==item.GoodsId{
				tmpMap:=map[string]interface{}{}
				tmpMap["id"]=item.Id
				tmpMap["goods_id"]=item.GoodsId
				tmpMap["goods_name"]=good.Name
				tmpMap["goods_image"]=good.GoodsFrontImage
				tmpMap["goods_price"]=good.ShopPrice
				tmpMap["nums"]=item.Nums
				tmpMap["checked"]=item.Checked
				goodsList=append(goodsList,tmpMap)

			}
		}
	}
	reMap:=gin.H{
		"total":rsp.Total,
	}
	reMap["data"]=goodsList
	c.JSON(http.StatusOK,reMap)

	/*
	{

	}
	*/


}
