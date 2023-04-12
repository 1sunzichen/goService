package global

import (
	"gopro/gin_test/order_srv/config"
	"gopro/gin_test/order_srv/proto"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
    err error
	ServerConfig config.ServerConfig
	GoodSrvClient proto.GoodsClient
	InventorySrvClient proto.InventoryClient
)

