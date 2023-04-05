package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gopro/gin_test/mxshop-api/user-web/global"
	"gopro/gin_test/mxshop-api/user-web/proto"

	//"gopro/gin_test/mxshop-api/user-web/proto"
)
func InitSrvConn(){
	 	consulInfo:=global.ServerConfig.ConsulInfo
	 	userInfo:=global.ServerConfig.UserSrvInfo
		userConn,err:=grpc.Dial(
			fmt.Sprintf("consul://%s:%d/%s?wait=14s",consulInfo.Host,consulInfo.Port, userInfo.Name),
		grpc.WithInsecure(),grpc.WithDefaultServiceConfig(
			`{"loadBalaningPolicy":"round_robin"}`))
		if err!=nil{
			zap.S().Fatal("【initsrvconn】用户服务链接失败")
		}
		fmt.Println(userConn,consulInfo.Host,consulInfo.Port, userInfo.Name,"**************")
		userSrvClient:=proto.NewUserClient(userConn)

		global.UserSrvClient=userSrvClient
}
func InitConsul(){
	var address string
	var port int
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",global.ServerConfig.ConsulInfo.Host,8500)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	un:=global.ServerConfig.UserSrvInfo.Name
	fmt.Printf("un%s",un)
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"",un))
	//_, err = client.Agent().ServicesWithFilter(`Service == "user_srv"`)
	if err != nil {
		panic(err)
	}
	//
	for key, value := range data {
		fmt.Println(key,value,value.Address,value.Port)
		address=value.Address
		port=value.Port
		break
	}
	userConn,err:=grpc.Dial(fmt.Sprintf("%s:%d",address,port ),grpc.WithInsecure())
	if err!=nil{
		zap.S().Errorw("[getuserlist] 链接【用户服务失败】",
			"msg",err.Error())
	}
	userSrvClient:=proto.NewUserClient(userConn)

	global.UserSrvClient=userSrvClient
}