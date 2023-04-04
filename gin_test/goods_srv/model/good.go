package model

// 类型,这个字段是否能为null 分类表结构
type Category struct{
	BaseModel
	Name string `gorm:"type:varchar(20);not null"`
	ParentCatgoryID int32
	//指向自己 指针
	ParentCatgory *Category
	//外健 ParentCatgoryID 指向那一列 ID
	SubCategory []*Category `gorm:"foreignKey:ParentCatgoryID;references:ID"`
	Level int32 `gorm:"type:int;not null;default:1"`
	IsTab bool `gorm:"default:false;not null"`
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
