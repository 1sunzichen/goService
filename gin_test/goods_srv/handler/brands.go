package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopro/gin_test/goods_srv/global"
	"gopro/gin_test/goods_srv/model"
	"gopro/gin_test/goods_srv/proto"
)

//品牌和轮播图
func (g *GoodsServer)BrandList(ctx context.Context,req *proto.BrandFilterRequest) (*proto.BrandListResponse, error){
	brandListResponse:=proto.BrandListResponse{}
	var brands []model.Brands
	//不分页
	//result:=global.DB.Find(&brands)
	//if result.Error!=nil{
	//	return nil, result.Error
	//}
	//分页
	result:=global.DB.Scopes(Paginate(int(req.Pages),int(req.PagePerNums))).Find(&brands)
	if result.Error!=nil{
		return nil, result.Error
	}
	var total int64
	global.DB.Model(&model.Brands{}).Count(&total)
	brandListResponse.Total=int32(total)
	//brandListResponse.Total=int32(result.RowsAffected)
	var brandResp []*proto.BrandInfoResponse
	for _,brand:=range  brands{
		brandResp=append(brandResp,&proto.BrandInfoResponse{
			Id:brand.ID,
			Name:brand.Name,
			Logo:brand.Logo,
		} )
	}
	brandListResponse.Data=brandResp
	//指针 去地址
	return &brandListResponse,nil
}
func (g *GoodsServer)CreateBrand(ctx context.Context,req *proto.BrandRequest) (*proto.BrandInfoResponse, error){
	var br map[string]interface{}
	if result:=global.DB.First(&model.Brands{},"name = ?",req.Name);result.RowsAffected==1{
		result.Scan(&br)
		fmt.Printf("传进来的名字%s,得到%s",req.Name,br["name"])

		return nil,status.Errorf(codes.InvalidArgument,"品牌已存在")
	}
	brand:=&model.Brands{
		Name: req.Name,
		Logo: req.Logo,
		BaseModel:model.BaseModel{
			IsDeleted:true,
		},
	}
	global.DB.Save(&brand)
	return &proto.BrandInfoResponse{Id:brand.ID},nil
}
func (g *GoodsServer)DeleteBrand(context.Context, *proto.BrandRequest) (*emptypb.Empty, error){
	if result:=global.DB.Delete(&model.Brands{});result.RowsAffected==0{
		return nil,status.Errorf(codes.NotFound,"品牌不存在")
	}

	return &emptypb.Empty{},nil
}

func (g *GoodsServer)UpdateBrand(ctx context.Context,req *proto.BrandRequest) (*emptypb.Empty, error){
	brand:=&model.Brands{}
	if result:=global.DB.First(&brand);result.RowsAffected==0{
		return nil,status.Errorf(codes.NotFound,"品牌不存在")
	}

	if req.Name!=""{
		brand.Name=req.Name
	}
	if req.Logo!=""{
		brand.Logo=req.Logo
	}
	global.DB.Save(&brand)
	return &emptypb.Empty{},nil
}
