package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopro/gin_test/mxshop-api/user-web/global"
	"gopro/gin_test/mxshop-api/user-web/initialize"
)

func main(){

	port:=8021
	initialize.InitLogger()
	initialize.InitConfig()
	r:=initialize.Routers()
	zap.S().Debugf("启动服务器%d",port)
	//if err:=r.Run(fmt.Sprintf(":%d",port));err!=nil{
	// zap.S().Panic("startup failed", err.Error())
	//}
	//Router:=gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200,"pong")
	})
	// TODO: issue here
	if err:=r.Run(fmt.Sprintf(":%d",global.ServerConfig.Port));err!=nil{
		zap.S().Panic("启动失败",err.Error())
	}
	//r.Run(":"+port)
}
