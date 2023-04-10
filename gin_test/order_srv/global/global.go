package global

import (
	"gopro/gin_test/inventory_srv/config"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
    err error
	ServerConfig config.ServerConfig
)

