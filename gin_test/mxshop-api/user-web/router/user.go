package router

import (
	"github.com/gin-gonic/gin"
	"gopro/gin_test/mxshop-api/user-web/api"
	"gopro/gin_test/mxshop-api/user-web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup){
	UserRouter:=Router.Group("user")
	{
		UserRouter.GET("list",middlewares.JWTAuth(),middlewares.IsAdminAuth(),api.GetUserList)
		UserRouter.POST("pwd_login",api.PasswordLogin)
		UserRouter.POST("register",api.Register)
	}

}