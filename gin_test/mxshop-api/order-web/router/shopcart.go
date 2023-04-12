package router

import (
	"github.com/gin-gonic/gin"
	"gopro/gin_test/mxshop-api/order-web/api/shop_cart"
	"gopro/gin_test/mxshop-api/order-web/middlewares"
)

func InitShopCartRouter(Router *gin.RouterGroup){
	ShopCartRouter:=Router.Group("shopcarts").Use(middlewares.JWTAuth())


	{
		ShopCartRouter.GET("",shop_cart.List)
		ShopCartRouter.DELETE("/:id",shop_cart.Delete)
		ShopCartRouter.PUT("/:id",shop_cart.Update)
		ShopCartRouter.POST("",shop_cart.New)
	}


}