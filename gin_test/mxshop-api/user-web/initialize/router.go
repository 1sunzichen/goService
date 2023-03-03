package initialize

import (
	"github.com/gin-gonic/gin"
	"gopro/gin_test/mxshop-api/user-web/router"
)

func Routers() *gin.Engine{
	Router:=gin.Default()
	Apigroup:=Router.Group("/u/v1")
	router.InitUserRouter(Apigroup)
	return Router
}

