package router

import (
	"github.com/gin-gonic/gin"
	"gopro/gin_test/mxshop-api/goods-web/api/goods"
)

func InitRouter(Router *gin.RouterGroup){
	GoodsRouter:=Router.Group("goods")
	{
		GoodsRouter.GET("",goods.List)
		//UserRouter.POST("pwd_login",api.PasswordLogin)
		//UserRouter.POST("register",api.Register)
	}

}