package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gopro/gin_test/inventory_srv/proto"
	"sync"
)

var Client proto.InventoryClient
var conn *grpc.ClientConn
func Init(){
	conn,err:=grpc.Dial("127.0.0.1:50051",grpc.WithInsecure())
	if err!=nil{
		panic(err)
	}
	//defer  conn.Close()
	Client=proto.NewInventoryClient(conn)
}
func TestSet(goodsid int32,stocks int32){
	rsp,err:=Client.SetInv(context.Background(),&proto.GoodsInvInfo{
		GoodsId: goodsid,
		Num: stocks,
	})
	if err!=nil{
		panic(err)
	}
	fmt.Println(rsp)
}
func TestGet(goodsid int32){
	rsp,err:=Client.InvDetail(context.Background(),&proto.GoodsInvInfo{
		GoodsId: goodsid,
	})
	if err!=nil{
		panic(err)
	}
	fmt.Println(rsp.Num)
}
func TestSellSync(wg *sync.WaitGroup){
	defer wg.Done()
	_,err:=Client.Sell(context.Background(),&proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{
			{GoodsId: 422, Num: 1},

		},
	})
	if err!=nil{
		panic(err)
	}

}
func TestSell(){
	_,err:=Client.Sell(context.Background(),&proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{
			{GoodsId: 421, Num: 1},

		},
	})
	if err!=nil{
		panic(err)
	}

}
func TestBack(){
	_,err:=Client.ReBack(context.Background(),&proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{
			{GoodsId: 11, Num: 10},
			{GoodsId: 11, Num: 20},
		},
	})
	if err!=nil{
		panic(err)
	}

}
func main(){
	Init()
	//TestSet(421,30)
	//TestGet(11)
	//TestSell()
	//TestBack()
	//var i int32
	//for i=421;i<844;i++{
	//	TestSet(i,30)
	//}
	//TestSet(421,10000)
	var wg sync.WaitGroup
	runLine:=10
	wg.Add(runLine)

	for i:=0;i<runLine;i++{
		go TestSellSync(&wg)
	}
	wg.Wait()
}