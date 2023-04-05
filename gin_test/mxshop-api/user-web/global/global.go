package global

import (
	ut "github.com/go-playground/universal-translator"
	"gopro/gin_test/mxshop-api/user-web/config"
	"gopro/gin_test/mxshop-api/user-web/proto"
)

var (
	Trans ut.Translator
	ServerConfig *config.ServerConfig=&config.ServerConfig{
		UserSrvInfo:config.UserSrvConfig{},
		AliSmsInfo: config.AliSmsConfig{
			ApiKey: "LTAI5tA8amT2P5TiYB73XTos",
			ApiSecrect: "W278E8lwasu1IrK0iZSKPlTKAnuHTe",
		},
		RedisInfo:config.RedisConfig{
			Port: 6379,
			Expire: 900,
		},
	}
	UserSrvClient proto.UserClient
)

var MobileValidator="mobile_validator"