package global

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var (
	DB *gorm.DB
    err error
)

func init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer (log output target, prefix, and log content -- translator's note)
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,   // Ignore ErrRecordNotFound error
			Colorful:      false,         // Disable colorful printing
		},
	)

	DB, err = gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/zc_shop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	//

}