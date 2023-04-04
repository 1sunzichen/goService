package handler

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopro/gin_test/goods_srv/proto"
)

//品牌分类
func (g *GoodsServer)CategoryBrandList(context.Context, *proto.CategoryBrandFilterRequest) (*proto.CategoryBrandListResponse, error){
	return nil,nil
}
//通过category获取brands
func (g *GoodsServer)GetCategoryBrandList(context.Context, *proto.CategoryInfoRequest) (*proto.BrandListResponse, error){
	return nil,nil
}
func (g *GoodsServer)CreateCategoryBrand(context.Context, *proto.CategoryBrandRequest) (*proto.CategoryBrandResponse, error){
	return nil,nil
}
func (g *GoodsServer)DeleteCategoryBrand(context.Context, *proto.CategoryBrandRequest) (*emptypb.Empty, error){
	return nil,nil
}
func (g *GoodsServer)UpdateCategoryBrand(context.Context, *proto.CategoryBrandRequest) (*emptypb.Empty, error){
	return nil,nil
}
