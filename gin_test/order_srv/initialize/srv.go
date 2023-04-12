package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gopro/gin_test/order_srv/global"
	"gopro/gin_test/order_srv/proto"
)

func InitSrvConn(){
	consulInfo:=global.ServerConfig.ConsulInfo
	goodsConn,err:=grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",consulInfo.Host,consulInfo.Port,global.ServerConfig.GoodsSrvConfig.Name),
	grpc.WithInsecure(),grpc.WithDefaultServiceConfig(
		`{"loadBalaningPolicy":"round_robin"}`))
	if err!=nil{
		zap.S().Fatal("【initsrvconn】用户服务链接失败")
	}
	goodsSrvClient:=proto.NewGoodsClient(goodsConn)

	global.GoodSrvClient=goodsSrvClient
	InventoryConn,err:=grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",consulInfo.Host,consulInfo.Port,global.ServerConfig.InventoryConfig.Name),
		grpc.WithInsecure(),grpc.WithDefaultServiceConfig(
			`{"loadBalaningPolicy":"round_robin"}`))
	if err!=nil{
		zap.S().Fatal("【initsrvconn】用户服务链接失败")
	}
	InventorySrvClient:=proto.NewInventoryClient(InventoryConn)

	global.InventorySrvClient=InventorySrvClient
}