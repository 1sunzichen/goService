package shop_cart

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopro/gin_test/mxshop-api/order-web/api"
	"gopro/gin_test/mxshop-api/order-web/forms"
	"gopro/gin_test/mxshop-api/order-web/global"
	"gopro/gin_test/mxshop-api/order-web/proto"
	"net/http"
	"strconv"
)

func List(c *gin.Context){
	//获取购物车商品
	userId,_:=c.Get("userId")

	fmt.Println(userId,"userId",c.Query("userId"))
	rsp,err:=global.OrderClient.CartItemList(context.Background(),&proto.UserInfo{

		Id:int32(userId.(uint)),
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

func New(c *gin.Context){
	itemForm:=forms.ShopCartItemForm{}
	if err:=c.ShouldBindJSON(&itemForm);err!=nil{
		api.HandleGrpcErrorToHttp(err,c)
		return
	}
	//为了严谨性 ，添加购物车之前，检查商品是否存在
	_,err:=global.GoodsClient.GetGoodsDetail(context.Background(),&proto.GoodInfoRequest{
		Id: itemForm.GoodsId,
	})
	if err!=nil{
		zap.S().Errorw("[list] 查询商品信息失败")
		api.HandleGrpcErrorToHttp(err,c)
		return
	}
	invRsp,err:=global.InventoryClient.InvDetail(context.Background(),&proto.GoodsInvInfo{
		GoodsId: itemForm.GoodsId,
	})
	if err!=nil{
		zap.S().Errorw("[list] 查询商品信息失败")
		api.HandleGrpcErrorToHttp(err,c)
		return
	}
	if invRsp.Num<itemForm.Nums{
		c.JSON(http.StatusBadRequest,gin.H{
			"nums":"库存不足",
		})
		return
	}
	//
	userId,_:=strconv.Atoi(c.Query("userId"))
	rsp,err:=global.OrderClient.CreateCartItem(context.Background(),&proto.CartItemRequest{
		GoodsId: itemForm.GoodsId,
		UserId: int32(userId),
		Nums: itemForm.Nums,
	})
	if err!=nil{
		zap.S().Errorw("添加到购物车失败")
		api.HandleGrpcErrorToHttp(err,c)
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"id":rsp.Id,
	})


}
func Delete(c *gin.Context){
	id,_:=strconv.Atoi(c.Param("id"))
	userId,_:=c.Get("userId")
	//userId,_:=strconv.Atoi(c.Query("userid"))
	fmt.Println(id,userId)
	_,err:=global.OrderClient.DeleteCartItem(context.Background(),&proto.CartItemRequest{
		//UserId: int32(userId),
		UserId: int32(userId.(uint)),
		GoodsId: int32(id),
	})
	if err!=nil{
		zap.S().Errorw("删除购物车记录失败")
		api.HandleGrpcErrorToHttp(err,c)
		return
	}
	c.Status(http.StatusOK)
}

func Update(c *gin.Context){
	goodsId,_:=strconv.Atoi(c.Param("id"))

	itemForm:=forms.ShopCartItemUpdateForm{}
	if err:=c.ShouldBindJSON(&itemForm);err!=nil{
		api.HandleValidatorError(c,err)
		return
	}
	userId,_:=c.Get("userId")
	request:=proto.CartItemRequest{
		UserId:     int32(userId.(uint)),
		GoodsId:    int32(goodsId) ,
		Nums:       itemForm.Nums,
		Checked:    false,
	}
	if itemForm.Checked!=nil{
		request.Checked=*itemForm.Checked
	}
	_,err:=global.OrderClient.UpdateCartItem(context.Background(),&request)
	if err!=nil{
		zap.S().Errorw("更新购物车记录失败")
		api.HandleGrpcErrorToHttp(err,c)
		return
	}
	c.Status(http.StatusOK)

}