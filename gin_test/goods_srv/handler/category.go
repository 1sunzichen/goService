package handler

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopro/gin_test/goods_srv/global"
	"gopro/gin_test/goods_srv/model"
	"gopro/gin_test/goods_srv/proto"
)

//商品分类
func (g *GoodsServer)GetAllCategorysList(context.Context, *emptypb.Empty) (*proto.CategoryListResponse, error){
	var Categories []model.Category
	global.DB.Preload("SubCategory").Find(&Categories)
	for _,category :=range Categories {
		fmt.Println(category.Name)
	}
	return nil,nil
}
//获取子分类
func (g *GoodsServer)GetSubCategory(context.Context, *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error){
	return nil,nil
}
func (g *GoodsServer)CreateCategory(context.Context, *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error){
	return nil,nil
}
func (g *GoodsServer)DeleteCategory(context.Context, *proto.DeleteCategoryRequest) (*emptypb.Empty, error){
	return nil,nil
}
func (g *GoodsServer)UpdateCategory(context.Context, *proto.CategoryInfoRequest) (*emptypb.Empty, error){
	return nil,nil
}