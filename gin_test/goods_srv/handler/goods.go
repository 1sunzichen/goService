package handler

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopro/gin_test/goods_srv/proto"

)


type GoodsServer struct{
	proto.UnimplementedGoodsServer
}

func  (g *GoodsServer)GoodsList(ctx context.Context, req *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error){
	return nil, nil
}
//现在用户提交订单有多个商品，你得批量查询商品的信息吧
func (g *GoodsServer)BatchGetGoods(context.Context, *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error){
	return nil, nil
}
func (g *GoodsServer)CreateGoods(context.Context, *proto.CreateGoodsInfo) (*proto.GoodsInfoResponse, error){
	return nil,nil
}
func (g *GoodsServer)DeleteGoods(context.Context, *proto.DeleteGoodsInfo) (*emptypb.Empty, error){
	return nil,nil
}
func (g *GoodsServer)UpdateGoods(context.Context, *proto.CreateGoodsInfo) (*emptypb.Empty, error){
	return nil,nil
}


