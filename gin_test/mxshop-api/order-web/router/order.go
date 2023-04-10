package router

import (
	"github.com/gin-gonic/gin"
	"gopro/gin_test/mxshop-api/order-web/api/order"
)

func InitRouter(Router *gin.RouterGroup){
	OrderRouter:=Router.Group("order")
	{
		OrderRouter.GET("",order.List)
	}


}