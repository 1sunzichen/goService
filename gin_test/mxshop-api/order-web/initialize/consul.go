package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gopro/gin_test/mxshop-api/order-web/global"
	"gopro/gin_test/mxshop-api/order-web/proto"

	//"gopro/gin_test/mxshop-api/order-web/proto"
)
func InitSrvConn(){
	 	consulInfo:=global.ServerConfig.ConsulInfo
	 	userInfo:=global.ServerConfig.GoodsSrvInfo
	 	orderInfo:=global.ServerConfig.OrderSrvInfo
		userConn,err:=grpc.Dial(
			fmt.Sprintf("consul://%s:%d/%s?wait=14s",consulInfo.Host,consulInfo.Port, userInfo.Name),
		grpc.WithInsecure(),grpc.WithDefaultServiceConfig(
			`{"loadBalaningPolicy":"round_robin"}`))
		if err!=nil{
			zap.S().Fatal("【initsrvconn】用户服务链接失败")
		}
		GoodsClient:=proto.NewGoodsClient(userConn)
		global.GoodsClient=GoodsClient
		orderConn,err:=grpc.Dial(
			fmt.Sprintf("consul://%s:%d/%s?wait=14s",consulInfo.Host,consulInfo.Port, orderInfo.Name),
			grpc.WithInsecure(),grpc.WithDefaultServiceConfig(
				`{"loadBalaningPolicy":"round_robin"}`))
		if err!=nil{
			zap.S().Fatal("【initsrvconn】用户服务链接失败")
		}
		OrderClient:=proto.NewOrderClient(orderConn)
		global.OrderClient=OrderClient
}