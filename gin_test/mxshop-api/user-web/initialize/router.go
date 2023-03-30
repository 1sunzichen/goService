package initialize

import (
	"github.com/gin-gonic/gin"
	"gopro/gin_test/mxshop-api/user-web/middlewares"
	"gopro/gin_test/mxshop-api/user-web/router"
)

func Routers() *gin.Engine{
	Router:=gin.Default()
	Router.Use(middlewares.Cors())
	Apigroup:=Router.Group("/u/v1")

	router.InitUserRouter(Apigroup)
	router.InitBaseRouter(Apigroup)

	return Router
}

