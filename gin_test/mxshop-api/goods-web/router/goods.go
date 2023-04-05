package router

import (
	"github.com/gin-gonic/gin"
	"gopro/gin_test/mxshop-api/goods-web/api/goods"
	"gopro/gin_test/mxshop-api/goods-web/middlewares"
)

func InitRouter(Router *gin.RouterGroup){
	GoodsRouter:=Router.Group("goods")
	{
		GoodsRouter.GET("",goods.List)
		GoodsRouter.POST("",middlewares.JWTAuth(),middlewares.IsAdminAuth(),goods.New)
		GoodsRouter.GET(":id",middlewares.JWTAuth(),middlewares.IsAdminAuth(),goods.Detail)

		GoodsRouter.DELETE("/:id",middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Delete) //删除商品
		GoodsRouter.GET("/:id/stocks", goods.Stocks) //获取商品的库存

		GoodsRouter.PUT("/:id",middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Update)
		GoodsRouter.PATCH("/:id",middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.UpdateStatus)
	}


}