package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gopro/gin_test/mxshop-api/goods-web/global"
	"gopro/gin_test/mxshop-api/goods-web/proto"

	//"gopro/gin_test/mxshop-api/goods-web/proto"
)
func InitSrvConn(){
	 	consulInfo:=global.ServerConfig.ConsulInfo
	 	userInfo:=global.ServerConfig.GoodsSrvInfo
		userConn,err:=grpc.Dial(
			fmt.Sprintf("consul://%s:%d/%s?wait=14s",consulInfo.Host,consulInfo.Port, userInfo.Name),
		grpc.WithInsecure(),grpc.WithDefaultServiceConfig(
			`{"loadBalaningPolicy":"round_robin"}`))
		if err!=nil{
			zap.S().Fatal("【initsrvconn】用户服务链接失败")
		}
		goodsSrvClient:=proto.NewGoodsClient(userConn)

		global.GoodsSrvClient=goodsSrvClient
}