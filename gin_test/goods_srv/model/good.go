package model

import (
	"context"
	"gopro/gin_test/goods_srv/global"
	"gorm.io/gorm"
	"strconv"
)

// 类型,这个字段是否能为null 分类表结构
type Category struct{
	BaseModel
	Name string `gorm:"type:varchar(20);not null" json:"name"`
	ParentCategoryID int32 `json:"parent"`
	//指向自己 指针
	ParentCategory *Category `json:"-"`
	//外健 ParentCatgoryID 指向那一列 ID
	SubCategory []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	Level int32 `gorm:"type:int;not null;default:1" json:"level"`
	IsTab bool `gorm:"default:false;not null" json:"is_tab"`
}

type Brands struct{
	BaseModel
	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(200);default:'';not null"`
}

type GoodsCategoryBrand struct{
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category Category
	Brands Brands
	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique"`

}

func (GoodsCategoryBrand) TableName() string{
	return "goodscategorybrand"
}

type Banner struct{
	BaseModel
	Image string `gorm:"type:varchar(200);not null"`
	Url string `gorm:"type:varchar(200);not null"`
	Index int32 `gorm:"type:int;default:1;not null"`
}

type Goods struct{
	BaseModel
	CategoryID int32 `gorm:"type:int;not null"`
	Category Category
	BrandsID int32 `gorm:"type:int;not null"`
	Brands Brands
	OnSale bool `gorm:"default:false;not null"`
	ShipFree bool `gorm:"default:false;not null"`
	IsNew bool `gorm:"default:false;not null"`
	IsHot bool `gorm:"default:false;not null"`
	Name string `gorm:"type:varchar(50);not null"`
	GoodsSn string `gorm:"type:varchar(50);not null"`
	ClickNum int32 `gorm:"type:int;default:0;not null"`
	SoldNum int32 `gorm:"type:int;default:0;not null"`
	FavNum int32 `gorm:"type:int;default:0;not null"`
	MarketPrice float32 `gorm:"not null"`
	ShopPrice float32 `gorm:"not null"`
	GoodsBrief string `gorm:"type:varchar"`
	Images GormList
	DescImages GormList
	GoodsFrontImage string `gorm:"type:varchar(200);not null"`
}
func (g *Goods)AfterCreate(tx *gorm.DB)(err error){
	err=CommonUpdateES(g,"")
	return err
}
func (g *Goods)AfterUpdate(tx *gorm.DB)(err error){
	err=CommonUpdateES(g,"update")
	return err
}

// CreateEsIndex 新增Es gorm 回调
func CreateEsIndex(esModel EsGoods,g *Goods)error{
	_,err:=global.EsClient.Index().Index("goods").BodyJson(esModel).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
	if err != nil {
		//panic(err)
		return err
	}
	return nil
}

// UpdateEsIndex 更新Es gorm 回调
func UpdateEsIndex(esModel EsGoods,g *Goods)error{
	_,err:=global.EsClient.Update().Index("goods").Doc(esModel).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
	if err != nil {
		//panic(err)
		return err
	}
	return nil
}

// DeleteEsIndex 删除Es gorm 回调
func DeleteEsIndex(esModel EsGoods,g *Goods)error{
	_,err:=global.EsClient.Delete().Index("goods").Id(strconv.Itoa(int(g.ID))).Do(context.Background())
	if err != nil {
		//panic(err)
		return err
	}
	return nil
}
func CommonUpdateES(g *Goods,key string) error {

	esModel := EsGoods{
		ID:          g.ID,
		CategoryID:  g.CategoryID,
		BrandsID:    g.BrandsID,
		OnSale:      g.OnSale,
		ShipFree:    g.ShipFree,
		IsNew:       g.IsNew,
		IsHot:       g.IsHot,
		Name:        g.Name,
		ClickNum:    g.ClickNum,
		SoldNum:     g.SoldNum,
		FavNum:      g.FavNum,
		MarketPrice: g.MarketPrice,
		GoodsBrief:  g.GoodsBrief,
		ShopPrice:   g.ShopPrice,
	}
	switch key {
		case "update":
			err :=UpdateEsIndex(esModel,g)
			if err != nil {
				//panic(err)
				return err
			}

		case "delete":
			err :=DeleteEsIndex(esModel,g)
			if err != nil {
				//panic(err)
				return err
			}
		default:
			err :=CreateEsIndex(esModel,g)
			if err != nil {
				//panic(err)
				return err
			}
	}


	//强调一下 一定要将docker启动es的java_ops的内存设置大一些 否则运行过程中会出现 bad request错误

	return nil
}