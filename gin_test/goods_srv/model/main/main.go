package main

import (
	//"crypto/sha512"
	//"fmt"
	//"github.com/anaskhan96/go-password-encoder"
	"gopro/gin_test/goods_srv/model"
	//"gopro/gin_test/goods_srv/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)



func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,   // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:      false,         // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/zc_shop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	_=db.AutoMigrate(&model.Category{},&model.Brands{},&model.GoodsCategoryBrand{},&model.Goods{})

	//options:=&password.Options{16,100,32,sha512.New}
	//salt,encodedPwd:=password.Encode("admin123",options)
	//newPassword:=fmt.Sprintf("$ekko307$%s$%s",salt,encodedPwd)
	//for i:=0;i<10;i++{
	//	user:=model.User{
	//		NickName: fmt.Sprintf("boddy%d",i),
	//		Mobile:fmt.Sprintf("boddy%d",i),
	//		Password:newPassword,
	//	}
	//	db.Save(&user)
	//}

	//
	//db.AutoMigrate(&models.User{})

}