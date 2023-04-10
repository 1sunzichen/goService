package initialize

import (
	"github.com/gin-gonic/gin"
	"gopro/gin_test/mxshop-api/order-web/middlewares"
	"gopro/gin_test/mxshop-api/order-web/router"
)

func Routers() *gin.Engine{
	Router:=gin.Default()
	Router.Use(middlewares.Cors())
	Apigroup:=Router.Group("/g/v1")

	router.InitRouter(Apigroup)
	router.InitShopCartRouter(Apigroup)
	return Router
}

