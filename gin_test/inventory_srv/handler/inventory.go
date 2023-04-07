package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopro/gin_test/inventory_srv/global"
	"gopro/gin_test/inventory_srv/model"
	"gopro/gin_test/inventory_srv/proto"
)

type Inventory struct{
	proto.UnimplementedInventoryServer
}
func (I *Inventory)SetInv(ctx context.Context,req *proto.GoodsInvInfo)(*empty.Empty,error){
	//设置库存 更新库存
	var inv model.Inventory
	global.DB.First(&inv,req.GoodsId)
	inv.Goods=req.GoodsId
	inv.Stocks=req.Num
	global.DB.Save(&inv)
	return &empty.Empty{},nil
}
func (I *Inventory)InvDetail(ctx context.Context,req *proto.GoodsInvInfo)(*proto.GoodsInvInfo,error){
	var inv model.Inventory
	if res:=global.DB.First(&inv,req.GoodsId);res.RowsAffected==0{
		return nil,status.Errorf(codes.NotFound,"没有库存信息")
	}
	return &proto.GoodsInvInfo{
		GoodsId: inv.Goods,
		Num: inv.Stocks,
	},nil
}

func (I *Inventory)Sell(ctx context.Context,req *proto.SellInfo)(*empty.Empty,error){
	//数据库事务
	//并发情况下 可能会出现超卖
	tx:=global.DB.Begin()
	for _,goodinfo:=range req.GoodsInfo{
		var inv model.Inventory
		if result :=global.DB.First(&inv,goodinfo.GoodsId);result.RowsAffected==0{
			tx.Rollback()
			return nil,status.Errorf(codes.InvalidArgument,"没有库存信息")
		}
		if inv.Stocks<goodinfo.Num{
			tx.Rollback()
			return nil,status.Errorf(codes.ResourceExhausted,"库存不足")
		}
		//扣减
		inv.Stocks-=goodinfo.Num
		tx.Save(&inv)
	}
	tx.Commit()
	return &empty.Empty{},nil
}
func (I *Inventory)ReBack(ctx context.Context,req *proto.SellInfo)(*empty.Empty,error){
	return nil,nil
}