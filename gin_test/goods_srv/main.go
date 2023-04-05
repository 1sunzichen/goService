package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"gopro/gin_test/goods_srv/global"
	"gopro/gin_test/goods_srv/handler"
	"gopro/gin_test/goods_srv/initialize"
	"gopro/gin_test/goods_srv/proto"
	"gopro/gin_test/goods_srv/utils"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	IP:=flag.String("ip","0.0.0.0","ip地址")
	Port:=flag.Int("port",0,"端口号")
	//初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	zap.S().Info(global.ServerConfig)
	flag.Parse()
	fmt.Printf("ip%s:,port%d:",*IP,*Port)
	if *Port==0{
		*Port,_=utils.GetFreePort()
	}
	client,serverId:=initialize.InitConsul(*Port)
	//
	zap.S().Info("port",Port)
	server:=grpc.NewServer()
	proto.RegisterGoodsServer(server,&handler.GoodsServer{})
	lis,err:=net.Listen("tcp",fmt.Sprintf("%s:%d",*IP,*Port))
	if err!=nil{
		panic("failed to listen:"+err.Error())
	}
	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server,health.NewServer())
	//err=server.Serve(lis)
	//if err!=nil{
	//	panic("启动grpc失败"+err.Error())
	//}
	//接受终止信号
	go func() {
		err=server.Serve(lis)
		if err!=nil{
			panic("启动grpc失败"+err.Error())
		}
	}()
	quit:=make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	if err=client.Agent().ServiceDeregister(serverId);err!=nil{
		zap.S().Info("注销失败")
	}
	zap.S().Info("注销成功")


}
