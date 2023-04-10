package handler

import (
	"context"
	"fmt"
	girdled "github.com/go-redis/redis/v7"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v7"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopro/gin_test/order_srv/global"
	"gopro/gin_test/order_srv/model"
	"gopro/gin_test/order_srv/proto"
	"sync"
)

type InventoryServer struct{
	proto.UnimplementedInventoryServer
}
func (I *InventoryServer)SetInv(ctx context.Context,req *proto.GoodsInvInfo)(*empty.Empty,error){
	//设置库存 更新库存
	var inv model.Inventory
	global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv)
	inv.Goods=req.GoodsId
	inv.Stocks=req.Num
	global.DB.Save(&inv)
	return &empty.Empty{},nil
}
func (I *InventoryServer)InvDetail(ctx context.Context,req *proto.GoodsInvInfo)(*proto.GoodsInvInfo,error){
	var inv model.Inventory
	if res:=global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv);res.RowsAffected==0{
		return nil,status.Errorf(codes.NotFound,"没有库存信息")
	}
	return &proto.GoodsInvInfo{
		GoodsId: inv.Goods,
		Num: inv.Stocks,
	},nil
}
var m sync.Mutex // 互斥锁

func Rs()*redsync.Redsync{
	client := girdled.NewClient(&girdled.Options{
		Addr: "127.0.0.1:6379",
	})
	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)

	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	rs := redsync.New(pool)
	return rs
}
func (I *InventoryServer)Sell(ctx context.Context,req *proto.SellInfo)(*empty.Empty,error) {
	//数据库事务
	//并发情况下 可能会出现超卖

	rs:=Rs()

	tx := global.DB.Begin()
	for _, goodinfo := range req.GoodsInfo {
		var inv model.Inventory
		mutex := rs.NewMutex(fmt.Sprintf("获取锁%d",goodinfo.GoodsId))
		if err := mutex.Lock(); err != nil {
			return nil, status.Errorf(codes.ResourceExhausted, "获取redis锁异常")
		}
		if result := global.DB.Where(&model.Inventory{Goods: goodinfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
				tx.Rollback()
				return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
			}
			if inv.Stocks < goodinfo.Num {
				tx.Rollback()
				return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
			}
			//扣减 会出现数据不一致的情况 -锁 分布式锁
			inv.Stocks -= goodinfo.Num
			tx.Save(&inv)
			if ok, err := mutex.Unlock(); !ok || err != nil {
				return nil, status.Errorf(codes.ResourceExhausted, "释放redis锁异常")

			}
	}
	tx.Commit()
	return &empty.Empty{}, nil
}
	//func (I *InventoryServer)Sell(ctx context.Context,req *proto.SellInfo)(*empty.Empty,error){
//	//数据库事务
//	//并发情况下 可能会出现超卖
//	tx:=global.DB.Begin()
//	for _,goodinfo:=range req.GoodsInfo{
//		var inv model.Inventory
//		for {
//			//if result :=tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where(&model.Inventory{Goods: goodinfo.GoodsId}).First(&inv);result.RowsAffected==0{
//			if result := global.DB.Where(&model.Inventory{Goods: goodinfo.GoodsId}).First(&inv); result.RowsAffected == 0 {
//				tx.Rollback()
//				return nil, status.Errorf(codes.InvalidArgument, "没有库存信息")
//			}
//			if inv.Stocks < goodinfo.Num {
//				tx.Rollback()
//				return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
//			}
//			//扣减 会出现数据不一致的情况 -锁 分布式锁
//			inv.Stocks -= goodinfo.Num
//			//tx.Save(&inv) //gorm 出现零值的情况
//			if result := tx.Model(&model.Inventory{}).Select("Stocks","Version").Where("goods = ? and version = ?", goodinfo.GoodsId, inv.Version).Updates(model.Inventory{
//				Version: inv.Version + 1,
//				Stocks:  inv.Stocks,
//			}); result.RowsAffected == 0 {
//				zap.S().Info("库存扣减失败")
//			} else {
//				break
//			}
//		}
//	}
//	tx.Commit()
//	return &empty.Empty{},nil
//}

func (I *InventoryServer)ReBack(ctx context.Context,req *proto.SellInfo)(*empty.Empty,error){
	//数据库事务
	//
	tx:=global.DB.Begin()
	//m.Lock()



	for _,goodinfo:=range req.GoodsInfo{
		var inv model.Inventory
		//if result :=global.DB.Where(&model.Inventory{Goods: goodinfo.GoodsId}).First(&inv);result.RowsAffected==0{
		//加入forupdate //乐观锁
		for {
		//if result :=tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where(&model.Inventory{Goods: goodinfo.GoodsId}).First(&inv);result.RowsAffected==0{
		if result :=global.DB.Where(&model.Inventory{Goods: goodinfo.GoodsId}).First(&inv);result.RowsAffected==0{
			tx.Rollback()
			return nil,status.Errorf(codes.InvalidArgument,"没有库存信息")
		}

		//扣减
		inv.Stocks+=goodinfo.Num

			if result:=tx.Model(&model.Inventory{}).Where("goods = ? and version = ?",goodinfo.GoodsId,inv.Version).Updates(model.Inventory{
				Version:inv.Version+1,
				Stocks: inv.Stocks,
			});result.RowsAffected==0{
				zap.S().Info("库存扣减失败")
			}else{
				break
			}
		//tx.Save(&inv)
		}
	}

	tx.Commit()
	//m.Unlock()
	return &empty.Empty{},nil
}