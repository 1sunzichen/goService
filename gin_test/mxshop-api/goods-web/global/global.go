package global

import (
	ut "github.com/go-playground/universal-translator"
	"gopro/gin_test/mxshop-api/goods-web/config"
	"gopro/gin_test/mxshop-api/goods-web/proto"
)

var (
	Trans ut.Translator
	ServerConfig *config.ServerConfig=&config.ServerConfig{

	}
	GoodsSrvClient proto.GoodsClient

)

var MobileValidator="mobile_validator"