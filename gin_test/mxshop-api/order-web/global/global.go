package global

import (
	ut "github.com/go-playground/universal-translator"
	"gopro/gin_test/mxshop-api/order-web/config"
	"gopro/gin_test/mxshop-api/order-web/proto"
)

var (
	Trans ut.Translator
	ServerConfig *config.ServerConfig=&config.ServerConfig{

	}
	OrderClient proto.OrderClient
	GoodsClient proto.GoodsClient
)

var MobileValidator="mobile_validator"