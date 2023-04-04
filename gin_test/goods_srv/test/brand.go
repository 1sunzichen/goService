package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gopro/gin_test/goods_srv/proto"
)

var Client proto.GoodsClient
var conn *grpc.ClientConn
func Init(){
	conn,err:=grpc.Dial("127.0.0.1:50051",grpc.WithInsecure())
	if err!=nil{
		panic(err)
	}
	//defer  conn.Close()
	Client=proto.NewGoodsClient(conn)
}
func TestGetList(){
	rsp,err:=Client.BrandList(context.Background(),&proto.BrandFilterRequest{

	})
	if err!=nil{
		panic(err)
	}
	for _,data:=range rsp.Data{
		fmt.Println(data.Name)


	}
}
func TestCreate(){
	//for i:=1;i<10;i++{
		createRsp,err:=Client.CreateBrand(context.Background(),&proto.BrandRequest{
			Name: "神秘文字",
			Logo: "abcdfgh",

		})
		if err!=nil{
			panic(err)
		}
		fmt.Println(createRsp.Id)
	//}
}
func main(){
	Init()
	//TestGetList()
	//conn.Close()
	//TestCreate()
	TestGetListCate()
}