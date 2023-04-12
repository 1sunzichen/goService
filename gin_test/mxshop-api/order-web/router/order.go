package router

import (
	"github.com/gin-gonic/gin"
	"gopro/gin_test/mxshop-api/order-web/api/order"
	"gopro/gin_test/mxshop-api/order-web/api/pay"
	"gopro/gin_test/mxshop-api/order-web/middlewares"
)

func InitRouter(Router *gin.RouterGroup){
	OrderRouter:=Router.Group("order").Use(middlewares.JWTAuth())
	{
		OrderRouter.GET("",order.List)
		OrderRouter.POST("",order.New)
		OrderRouter.GET("/:id",order.Detail)

	}
	PayRouter:=Router.Group("pay")
	{
		PayRouter.POST("alipay/notify",pay.Notify)
	}
}