package router

import (
	"github.com/gin-gonic/gin"
	"gopro/gin_test/mxshop-api/user-web/api"
)

func InitBaseRouter(Router *gin.RouterGroup){
	BaseRouter:=Router.Group("base")
	{
		BaseRouter.GET("captcha",api.GetCaptcha)
		BaseRouter.GET("sms",api.SendSms)
	}
}
