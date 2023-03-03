package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"gopro/gin_test/mxshop_srv/handler"
	"gopro/gin_test/mxshop_srv/proto"
	"net"
)

func main(){
	IP:=flag.String("ip","127.0.0.1","ip地址")
	Port:=flag.Int("port",50051,"端口号")
	flag.Parse()
	fmt.Printf("ip%s:,port%d:",*IP,*Port)
	server:=grpc.NewServer()
	proto.RegisterUserServer(server,&handler.UserServer{})
	lis,err:=net.Listen("tcp",fmt.Sprintf("%s:%d",*IP,*Port))
	if err!=nil{
		panic("failed to listen:"+err.Error())
	}
	err=server.Serve(lis)
	if err!=nil{
		panic("启动grpc失败"+err.Error())
	}


}
