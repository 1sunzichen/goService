package global

import (
	"gopro/gin_test/mxshop_srv/config"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
    err error
	ServerConfig config.ServerConfig
)

//func init() {
//	//newLogger := logger.New(
//	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
//	//	logger.Config{
//	//		SlowThreshold: time.Second,   // 慢 SQL 阈值
//	//		LogLevel:      logger.Silent, // 日志级别
//	//		IgnoreRecordNotFoundError: true,   // 忽略ErrRecordNotFound（记录未找到）错误
//	//		Colorful:      false,         // 禁用彩色打印
//	//	},
//	//)
//
//	//DB, err = gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/zc_shop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
//	//	Logger: newLogger,
//	//	NamingStrategy: schema.NamingStrategy{
//	//		SingularTable: true,
//	//	},
//	//})
//	//if err != nil {
//	//	panic("failed to connect database")
//	//}
//
//	//
//
//}