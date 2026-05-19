package main

import (
	"gopro/gin_test/mxshop_srv/model"
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
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer (log output target, prefix, and log content -- translator's note)
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,   // Ignore ErrRecordNotFound error
			Colorful:      false,         // Disable colorful printing
		},
	)

	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/zc_shop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	//
	db.AutoMigrate(&model.User{})

}