package order

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"gopro/gin_test/mxshop-api/order-web/api"
	"gopro/gin_test/mxshop-api/order-web/api/pay"
	"gopro/gin_test/mxshop-api/order-web/forms"
	"gopro/gin_test/mxshop-api/order-web/global"
	"gopro/gin_test/mxshop-api/order-web/models"
	"gopro/gin_test/mxshop-api/order-web/proto"
	"net/http"
	"strconv"
)

func List(c *gin.Context){
	userId,_:=c.Get("userId")
	claims,_:=c.Get("claims")
	request:=proto.OrderFilterRequest{}
	model:=claims.(*models.CustomClaims)
	if model.AuthorityId==1{
		request.UserId=int32(userId.(uint))
	}
	pages:=c.DefaultQuery("p","0")
	pageInt,_:=strconv.Atoi(pages)
	request.Pages=int32(pageInt)

	nums:=c.DefaultQuery("pnum","0")
	pNums,_:=strconv.Atoi(nums)
	request.PagePerNums=int32(pNums)
	fmt.Println(request.UserId,"userIdrequest")
	rsp,err:=global.OrderClient.OrderList(context.Background(),&request)
	if err!=nil{
		zap.S().Errorw("[List]查询 【订单列表】失败")
		api.HandleGrpcErrorToHttp(err,c)
		return
	}

	reMap:=gin.H{
		"total":rsp.Total,
		//"data":rsp.Data,
	}
	orderList:=make([]interface{},0)
	for _,item:=range rsp.Data{
		tmpMap:=map[string]interface{}{}
		tmpMap["id"]=item.Id
		tmpMap["status"]=item.Status
		tmpMap["user"]=item.UserId
		tmpMap["total"]=item.Total
		//tmpMap["id"]=item.id
		tmpMap["order_sn"]=item.OrderSn
		tmpMap["add_time"]=item.AddTime
		orderList=append(orderList,tmpMap)
	}

	reMap["data"]=orderList
	c.JSON(200,reMap)
}
func New(c *gin.Context){
	orderForm:=forms.CreateOrderForm{}
	if err:=c.ShouldBindJSON(&orderForm);err!=nil{
		api.HandleValidatorError(c,err)
	}
	userId,_:=c.Get("userId")
	fmt.Println(userId,"userId####")
	rsp,err:=global.OrderClient.Create(context.Background(),&proto.OrderRequest{
		UserId: int32(userId.(uint)),
		Name:orderForm.Name,
		Mobile: orderForm.Mobile,
		Address: orderForm.Address,
		Post: orderForm.Post,
	})
	if err!=nil{
		s, _ := status.FromError(err)
		zap.L().Info("订单服务出错:订单服务消息信息 "+s.Message())
		//zap.S().Errorw("订单服务出错:订单服务消息信息 ",s.Message())

		api.HandleGrpcErrorToHttp(err,c)
		return
	}

	// 支付宝url
	url,err:=pay.GenPayUrl(rsp)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":err.Error(),
		})
	}
	c.JSON(200,gin.H{
		"id":rsp.Id,
		"alipay_url":url,
	})
}
func SetUserID(c *gin.Context,op func(int32)){
	userId,_:=c.Get("userId")
	claims,_:=c.Get("claims")
	model:=claims.(*models.CustomClaims)
	if model.AuthorityId==1{
		op(int32(userId.(uint)))
	}
}
func Detail(c *gin.Context){
	id,_:=strconv.Atoi(c.Param("id"))
	request:=proto.OrderRequest{Id: int32(id)}
	SetUserID(c, func(i int32) {
		request.UserId=i
	})
	rsp,err:=global.OrderClient.OrderDetail(context.Background(),&request)
	if err!=nil{
		zap.S().Errorw("获取订单失败")
		api.HandleGrpcErrorToHttp(err,c)
	}
	reMap:=gin.H{}
	reMap["id"]=rsp.OrderInfo.Id
	goodsList:=make([]interface{},0)
	for _,item:=range rsp.Goods{
		tmpMap:=gin.H{
			"id":item.GoodsId,
			"name":item.GoodsName,
			"image":item.GoodsImage,
			"price":item.GoodsPrice,
			"nums":item.Nums,
		}
		goodsList=append(goodsList,tmpMap)
	}
	reMap["goods"]=goodsList
	// 支付宝url
	url,err:=pay.GenPayUrl(rsp.OrderInfo)

	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":err.Error(),
		})
	}
	reMap["alipay_url"]=url
	c.JSON(200,reMap)
}