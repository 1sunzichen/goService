package router

import (
	"github.com/gin-gonic/gin"
	"gopro/gin_test/mxshop-api/user-web/api"
)

func InitUserRouter(Router *gin.RouterGroup){
	UserRouter:=Router.Group("user")
	{
		UserRouter.GET("list",api.GetUserList)
	}

}