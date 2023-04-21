package main

import (
	"context"
	"github.com/olivere/elastic/v7"
	"gopro/gin_test/goods_srv/global"
	"strconv"

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
	// 全局模式

	host := "http://192.168.155.202:9200"
	logger := log.New(os.Stdout, "mxshop", log.LstdFlags)
	global.EsClient, err = elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false),
		elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}

	var goods []model.Goods
	db.Find(&goods)
	for _, g := range goods {
		esModel := model.EsGoods{
			ID:          g.ID,
			CategoryID:  g.CategoryID,
			BrandsID:    g.BrandsID,
			OnSale:      g.OnSale,
			ShipFree:    g.ShipFree,
			IsNew:       g.IsNew,
			IsHot:       g.IsHot,
			Name:        g.Name,
			ClickNum:    g.ClickNum,
			SoldNum:     g.SoldNum,
			FavNum:      g.FavNum,
			MarketPrice: g.MarketPrice,
			GoodsBrief:  g.GoodsBrief,
			ShopPrice:   g.ShopPrice,
		}

		_, err = global.EsClient.Index().Index("goods").BodyJson(esModel).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
		if err != nil {
			panic(err)
		}
		//强调一下 一定要将docker启动es的java_ops的内存设置大一些 否则运行过程中会出现 bad request错误
	}
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//_=db.AutoMigrate(&model.Category{},&model.Brands{},&model.GoodsCategoryBrand{},&model.Goods{})

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
