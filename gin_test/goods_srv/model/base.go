package model

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)
type GormList []string
func (g GormList) Value()(driver.Value,error){
	return json.Marshal(g)
}

func (g *GormList)Scan(value interface{})error{
	return json.Unmarshal(value.([]byte),&g)
}
type BaseModel struct{
	ID int32 `gorm:"primarykey;type:int"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt `json:"-"`
	IsDeleted bool
}