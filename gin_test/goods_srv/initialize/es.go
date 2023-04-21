package initialize

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gopro/gin_test/goods_srv/global"
	"gopro/gin_test/goods_srv/model"
	"log"
	"os"
)

func InitEs ()  {
	esInfo:=global.ServerConfig.EsInfo
	host:=fmt.Sprintf("http://%s:%d",esInfo.Host,esInfo.Port)

	logger:=log.New(os.Stdout,"zc_shop",log.LstdFlags)

	client,err:=elastic.NewClient(elastic.SetURL(host),elastic.SetSniff(false),
		elastic.SetTraceLog(logger))
	if err!=nil{
		panic(err)
	}

	global.EsClient=client

	exists,err:=global.EsClient.IndexExists("goods").Do(context.Background())
	if err!=nil{
		panic(err)
	}
	if exists{
		_, err = global.EsClient.DeleteIndex("goods").Do(context.Background())
		if err != nil {
			log.Fatalf("Error deleting the index: %s", err)
		}
	}

	_,err=global.EsClient.CreateIndex("goods").BodyString(model.EsGoods{}.GetMapping()).Do(context.Background())
	if err!=nil{
		panic(err)
	}

	go func() {
		//初始化数据
		InitESData()
		//
	}()
}
func InitESData()error{
	var goods []model.Goods
	global.DB.Find(&goods)
	for _, g := range goods {

		err := model.CommonUpdateES(&g,"")
		if err != nil {
			//panic(err)
			return err
		}
		//强调一下 一定要将docker启动es的java_ops的内存设置大一些 否则运行过程中会出现 bad request错误
	}
	return nil
}