package main

import (
	"flag"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gopro/gin_test/inventory_srv/global"
	"gopro/gin_test/inventory_srv/handler"
	"gopro/gin_test/inventory_srv/initialize"
	"gopro/gin_test/inventory_srv/proto"
	"gopro/gin_test/inventory_srv/utils"
	"gopro/gin_test/inventory_srv/utils/register/consul"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	IP:=flag.String("ip","0.0.0.0","ip地址")
	Port:=flag.Int("port",50051,"端口号")
	//初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	//链接redis分布式锁

	//监听端口号
	zap.S().Info(global.ServerConfig)
	flag.Parse()
	fmt.Printf("ip%s:,port%d:",*IP,*Port)
	if *Port==0{
		*Port,_=utils.GetFreePort()
	}
	zap.S().Info("port",Port)
	//启动服务
	server:=grpc.NewServer()
	proto.RegisterInventoryServer(server,&handler.InventoryServer{})
	lis,err:=net.Listen("tcp",fmt.Sprintf("%s:%d",*IP,*Port))
	go func() {
		err=server.Serve(lis)
		if err!=nil{
			panic("启动grpc失败"+err.Error())
		}
	}()
	// 注册consul服务
	registerClient :=consul.NewRegister(global.ServerConfig.ConsulInfo.Host,global.ServerConfig.ConsulInfo.Port)
	serviceId:=fmt.Sprintf("%s",uuid.NewV4())
	err=registerClient.Register("127.0.0.1",*Port,global.ServerConfig.Name,[]string{"service服务","库存服务"},serviceId)

	if err!=nil{
		panic("failed to listen:"+err.Error())
	}
	//注册服务健康检查
	//grpc_health_v1.RegisterHealthServer(server,health.NewServer())

	//接受终止信号
	quit:=make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	if err=registerClient.DeRegister(serviceId);err!=nil{
		zap.S().Info("服务发现注销inventory_srv失败")
	}else{
		zap.S().Info("服务发现注销inventory_srv成功")
	}
}
