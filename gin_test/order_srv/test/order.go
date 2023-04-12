package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gopro/gin_test/order_srv/proto"
)

var Client proto.OrderClient
var conn *grpc.ClientConn
func Init(){
	conn,err:=grpc.Dial("127.0.0.1:56201",grpc.WithInsecure())
	if err!=nil{
		panic(err)
	}
	//defer  conn.Close()
	Client=proto.NewOrderClient(conn)
}
func TestCreate(UserId int32,GoodsId int32){
	rsp,err:=Client.CreateCartItem(context.Background(),&proto.CartItemRequest{

		UserId:     UserId,
		GoodsId:    GoodsId,
		Nums:       1,
		Checked: true,
	})
	if err!=nil{
		panic(err)
	}
	fmt.Println(rsp.Id)
}
func TestGetList(){
	rsp,err:=Client.CartItemList(context.Background(),&proto.UserInfo{Id: 1})
	if err!=nil{
		panic(err)
	}
	for _,item:=range rsp.Data{
		fmt.Println(item.Id,item.Nums,item.GoodsId)
	}
}

func TestCreateOrder(UserId int32){
	_,err:=Client.Create(context.Background(),&proto.OrderRequest{
		UserId:  UserId,
		Address: "佰嘉城",
		Name:    "zc",
		Mobile:  "15210187668",
		Post:    "快点发",

	})

	if err!=nil{
		panic(err)
	}
}
func main(){
	Init()
	var UserId int32=1
	TestCreate(UserId,421)
	//TestCreateOrder(UserId)
}