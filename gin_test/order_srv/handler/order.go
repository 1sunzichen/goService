package handler

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopro/gin_test/order_srv/global"
	"gopro/gin_test/order_srv/model"
	"gopro/gin_test/order_srv/proto"
	"math/rand"
	"time"
)

type OrderServer struct {
	*proto.UnimplementedOrderServer

}
func GenOrderSn(userId int32)string{
	now:=time.Now()
	rand.Seed(time.Now().UnixNano())
	orderSn:=fmt.Sprintf("%d%d%d%d%d%d%d%d",
		now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Nanosecond(),
		userId,rand.Intn(90)+10)
	return orderSn

}
func(*OrderServer)CartItemList(ctx context.Context,req *proto.UserInfo) (*proto.CartItemListResponse,error){
	var shopCarts []model.ShoppingCart
	var rsp proto.CartItemListResponse
	if result:=global.DB.Where(&model.ShoppingCart{User:req.Id}).Find(&shopCarts);result.Error!=nil{
		return nil,result.Error
	}else{
		rsp.Total=int32(result.RowsAffected)
	}
	for _,shopCart:=range shopCarts{
		rsp.Data=append(rsp.Data,&proto.ShopCartInfoResponse{
			Id: shopCart.ID,
			UserId: shopCart.User,
			GoodsId: shopCart.Goods,
			Nums:shopCart.Nums,
			Checked: shopCart.Checked,
		})
	}
	return &rsp,nil
}

func(*OrderServer)CreateCartItem(ctx context.Context,req *proto.CartItemRequest)(*proto.ShopCartInfoResponse,error){
	var shopCart model.ShoppingCart
	//查询
	if result:=global.DB.Where(&model.ShoppingCart{Goods: req.GoodsId,User: req.UserId}).First(&shopCart);result.RowsAffected!=0{
		shopCart.Nums+=req.Nums
	}else{
		//插入
		shopCart.User=req.UserId
		shopCart.Goods=req.GoodsId
		shopCart.Checked=req.Checked||false
		shopCart.Nums=req.Nums
	}
	global.DB.Save(&shopCart)
	return &proto.ShopCartInfoResponse{Id: shopCart.ID},nil
}


func(*OrderServer)UpdateCartItem(ctx context.Context,req *proto.CartItemRequest)(*empty.Empty,error){
	//更新购物车记录，更新数量和选中状态
	var shopCart model.ShoppingCart
	if result:=global.DB.Where("goods=? and user=?",req.GoodsId,req.UserId).First(&shopCart);result.RowsAffected==0{
		return nil,status.Errorf(codes.InvalidArgument,"记录不存在")
	}
	shopCart.Checked=req.Checked
	if req.Nums>0{
		shopCart.Nums=req.Nums
	}
	global.DB.Save(&shopCart)
	return &empty.Empty{},nil
}
func(*OrderServer)DeleteCartItem(ctx context.Context,req *proto.CartItemRequest)(*empty.Empty,error){
	//
	//删除购物车记录
	if result:=global.DB.Where("goods=? and user=?",req.GoodsId,req.UserId).Delete(&model.ShoppingCart{});result.RowsAffected==0{
		return nil,status.Errorf(codes.NotFound,"记录不存在")
	}

	return &empty.Empty{},nil
}
func(*OrderServer)OrderList(ctx context.Context,req *proto.OrderFilterRequest)(*proto.OrderListResponse,error){

	var rsp proto.OrderListResponse
	var orders []model.OrderInfo
	var total int64
	global.DB.Model(&model.OrderInfo{}).Where("user=?",req.UserId).Count(&total)
	fmt.Println(total,"Total",req.UserId)
	rsp.Total=int32(total)
	global.DB.Scopes(Paginate(int(req.Pages),int(req.PagePerNums))).Where(&model.OrderInfo{User: req.UserId}).Find(&orders)
	for _,order:=range orders{
		rsp.Data=append(rsp.Data,&proto.OrderInfoResponse{
			Id:      order.ID,
			UserId:  order.User,
			OrderSn: order.OrderSn,
			PayType: order.PayType,
			Status:  order.Status,
			Post:    order.Post,
			Total:   order.OrderMount,
			Address: order.Address,
			Name:    order.SignerName,
			Mobile:  order.SingerMobile,
			AddTime: order.CreatedAt.Format("2006-01-02"),
		})
	}
	return &rsp,nil
}
func(*OrderServer)OrderDetail(ctx context.Context,req *proto.OrderRequest)(*proto.OrderInfoDetailResponse,error){
	var rsp proto.OrderInfoDetailResponse
	//1
	var order model.OrderInfo
	//这个订单的id是否是当前用户的id
	//后台管理系统 orderid 电商系统 用户id  userid
	if result:=global.DB.Where(&model.OrderInfo{BaseModel:model.BaseModel{ID: req.Id}}).First(&order);result.Error!=nil{
		return nil,status.Errorf(codes.NotFound,"订单不存在")
	}
	orderInfo:=proto.OrderInfoResponse{}
	orderInfo.OrderSn=order.OrderSn
	orderInfo.Id=order.ID
	orderInfo.Address=order.Address
	orderInfo.Post=order.Post
	orderInfo.UserId=order.User
	orderInfo.Status=order.Status
	orderInfo.PayType=order.PayType
	orderInfo.Total=order.OrderMount
	orderInfo.Name=order.SignerName
	orderInfo.Mobile=order.SingerMobile
	rsp.OrderInfo=&orderInfo
	var orderGoods []model.OrderGoods
	if result:=global.DB.Where(&model.OrderGoods{Order: order.ID}).Find(&orderGoods);result.Error!=nil{
		return nil,status.Errorf(codes.NotFound,"商品信息不存在")
	}
	for _,orderGood :=range orderGoods{
		rsp.Goods=append(rsp.Goods,&proto.OrderItemResponse{
			GoodsId: orderGood.Goods,
			OrderId: orderGood.Order,
			Nums: orderGood.Nums,
			Id: orderGood.ID,
			GoodsImage: orderGood.GoodsImage,
			GoodsPrice: orderGood.GoodsPrice,
			GoodsName: orderGood.GoodsName,
		})
	}
	return &rsp,nil
}
func(*OrderServer)Create(ctx context.Context,req *proto.OrderRequest)(*proto.OrderInfoResponse,error){
	var goodIds []int32
	var shopCarts []model.ShoppingCart

	if result:=global.DB.Where(&model.ShoppingCart{User: req.UserId,Checked: true}).Find(&shopCarts);result.RowsAffected==0{
		return nil,status.Errorf(codes.InvalidArgument,"没有结算选中的商品")
	}
	goodsNumsMap:=make(map[int32]int32)
	for _,shopCart:=range shopCarts{
		goodIds=append(goodIds,shopCart.Goods)
		goodsNumsMap[shopCart.Goods]=shopCart.Nums
	}
	//

	goods,err:=global.GoodSrvClient.BatchGetGoods(context.Background(),&proto.BatchGoodsIdInfo{Id: goodIds})
	if err!=nil{
		return nil,status.Errorf(codes.Internal,"批量查询商品失败")
	}
	//
	var orderAmount float32
	var orderGoods []*model.OrderGoods
	var goodsInfo []*proto.GoodsInvInfo
	for _,good:=range goods.Data{
		orderAmount+=good.ShopPrice*float32(goodsNumsMap[good.Id])
		orderGoods=append(orderGoods,&model.OrderGoods{

			Goods:      good.Id,
			GoodsName:  good.Name,
			GoodsImage: good.GoodsFrontImage,
			GoodsPrice: good.ShopPrice,
			Nums:       goodsNumsMap[good.Id],
		})
		goodsInfo=append(goodsInfo,&proto.GoodsInvInfo{
			GoodsId: good.Id,
			Num: goodsNumsMap[good.Id],
		})
	}
	//跨服务调用库存服务进行扣减

	if _,err=global.InventorySrvClient.Sell(context.Background(),&proto.SellInfo{GoodsInfo: goodsInfo});err!=nil{
		return nil,status.Errorf(codes.ResourceExhausted,"库存扣减失败")

	}
	tx:=global.DB.Begin()
	//生成订单表
	order:=model.OrderInfo{
		OrderSn:      GenOrderSn(req.UserId),
		OrderMount:   orderAmount,
		Address:      req.Address,
		SignerName:   req.Name,
		SingerMobile: req.Mobile,
		Post:         req.Post,
		User: req.UserId,
	}
	fmt.Println(GenOrderSn(req.UserId),orderAmount)
	if result:=tx.Save(&order);result.RowsAffected==0{
		tx.Rollback()
		return nil,status.Errorf(codes.Internal,"创建订单失败")
	}

	for _,orderGood:=range orderGoods{
		orderGood.Order=order.ID
	}
	//
	if result:=tx.CreateInBatches(orderGoods,100);result.RowsAffected==0{
		tx.Rollback()
		return nil,status.Errorf(codes.Internal,"创建订单失败")

	}
	if result:=tx.Where(&model.ShoppingCart{User: req.UserId,Checked: true}).Delete(&model.ShoppingCart{});result.RowsAffected==0{

		tx.Rollback()
		return nil,status.Errorf(codes.Internal,"创建订单失败")

	}
	tx.Commit()
	return &proto.OrderInfoResponse{Id: order.ID,OrderSn: order.OrderSn,Total: order.OrderMount},nil
}
func(*OrderServer)UpdateOrderStatus(ctx context.Context,req *proto.OrderStatus) (*empty.Empty,error){
	if result:=global.DB.Model(&model.OrderInfo{}).Where("order_sn=?",req.OrderSn).Update("status",req.Status);result.RowsAffected==0{
		return nil,status.Errorf(codes.NotFound,"订单不存在")
	}
	return &empty.Empty{},nil
}