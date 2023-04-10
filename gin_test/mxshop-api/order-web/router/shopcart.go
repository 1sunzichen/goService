package router

import (
	"github.com/gin-gonic/gin"
	"gopro/gin_test/mxshop-api/order-web/api/shop_cart"
)

func InitShopCartRouter(Router *gin.RouterGroup){
	ShopCartRouter:=Router.Group("shopcarts")


	{
		ShopCartRouter.GET("",shop_cart.List)
	}


}