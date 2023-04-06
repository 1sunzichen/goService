package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
)



func TestGetListCate(){
	rsp,err:=Client.GetAllCategorysList(context.Background(),&empty.Empty{

	})
	if err!=nil{
		panic(err)
	}
	for _,data:=range rsp.Data{
		fmt.Println(data.Name)


	}
}
//func TestCreate(){
//	//for i:=1;i<10;i++{
//		createRsp,err:=Client.CreateBrand(context.Background(),&proto.BrandRequest{
//			Name: "神秘文字",
//			Logo: "abcdfgh",
//
//		})
//		if err!=nil{
//			panic(err)
//		}
//		fmt.Println(createRsp.Id)
//	//}
//}
//func main(){
//	Init()
//	TestGetListCate()
//	//conn.Close()
//	//TestCreate()
//}