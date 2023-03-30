package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gopro/gin_test/mxshop_srv/proto"
)

var userClient proto.UserClient
var conn *grpc.ClientConn
func Init(){
	conn,err:=grpc.Dial("127.0.0.1:50051",grpc.WithInsecure())
	if err!=nil{
		panic(err)
	}
	//defer  conn.Close()
	userClient=proto.NewUserClient(conn)
}
func TestGetUserList(){
	rsp,err:=userClient.GetUserList(context.Background(),&proto.PageInfo{
		Pn: 1,
		PSize: 2,
	})
	if err!=nil{
		panic(err)
	}
	for _,user:=range rsp.Data{
		fmt.Println(user.Mobile,user.Nickname,user.PassWord)
		checkRsp,err:=userClient.CheckPassWord(context.Background(),&proto.PassWordInfo{
			PassWord: "admin123",
			EncryptedPassword: user.PassWord,
		})
		if err!=nil{
			panic(err)
		}
		fmt.Println(checkRsp.Success)

	}
}
func TestCreateUser(){
	for i:=1;i<10;i++{
		createRsp,err:=userClient.CreateUser(context.Background(),&proto.CreateUserInfo{
			NickName: fmt.Sprintf("boddy%d",i),
			Mobile:fmt.Sprintf("boddy%d",i),
			PassWord:"admin123",
		})
		if err!=nil{
			panic(err)
		}
		fmt.Println(createRsp.Id)
	}
}
func main(){
	Init()
	TestCreateUser()
	//TestGetUserList()
	conn.Close()
}