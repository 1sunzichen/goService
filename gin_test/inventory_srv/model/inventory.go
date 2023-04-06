package model

type Stock struct{
	BaseModel
}
type Inventory struct{
	BaseModel
	Goods int32 `gorm:"type:int;index"`
	Stocks int32 `gorm:"type:int"`
	Version int32 `gorm:"type:int"` //分布式乐观锁
}
