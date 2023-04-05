package goods

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopro/gin_test/mxshop-api/goods-web/global"
	"gopro/gin_test/mxshop-api/goods-web/proto"
	"strconv"

	//"gopro/gin_test/mxshop-api/goods-web/proto"

	"net/http"
	"strings"
)
func removeTopStruct(fields map[string]string) map[string]string{
	rsp:=map[string]string{}
	for field,err:=range fields{
		rsp[field[strings.Index(field,".")+1:]]=err
	}
	return rsp
}
func HandleGrpcErrorToHttp(err error,c *gin.Context){
	//将grpc 的code转换成http的状态玛
	if err!=nil{
		if e,ok:=status.FromError(err);ok{
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})

			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": "参数",
				})
			default:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": "其他",
				})
			}
			return

		}
	}

}

func HandleVaildatorError(c *gin.Context,err error){
	errs,ok:=err.(validator.ValidationErrors)
	if !ok{
		c.JSON(http.StatusOK,gin.H{
			"msg":err.Error(),
		})
	}
	fmt.Println(errs,"错误")
	c.JSON(http.StatusBadRequest,gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}
func GetConsul()(string,int){
	var address string
	var port int
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",global.ServerConfig.ConsulInfo.Host,8500)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
    un:=global.ServerConfig.GoodsSrvInfo.Name
    fmt.Printf("un%s",un)
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"",un))
	//_, err = client.Agent().ServicesWithFilter(`Service == "goods_srv"`)
	if err != nil {
		panic(err)
	}
	//
 	for key, value := range data {
		fmt.Println(key,value,value.Address,value.Port)
		address=value.Address
		port=value.Port
		//break
	}
	return  address,port
}
func List(c *gin.Context)  {
	//拼装请求
	request:=&proto.GoodsFilterRequest{}
	//*********
	priceMin:=c.DefaultQuery("pmin","0")
	priceMinInt,_:=strconv.Atoi(priceMin)
	request.PriceMin=int32(priceMinInt)

	priceMax:=c.DefaultQuery("pmax","0")
	priceMaxInt,_:=strconv.Atoi(priceMax)
	request.PriceMax=int32(priceMaxInt)

	categoryId:=c.DefaultQuery("c","0")
	categoryIdInt,_:=strconv.Atoi(categoryId)
	request.TopCategory=int32(categoryIdInt)

	pages:=c.DefaultQuery("p","0")
	pagesInt,_:=strconv.Atoi(pages)
	request.Pages=int32(pagesInt)

	pageNum:=c.DefaultQuery("pnum","0")
	pageNumInt,_:=strconv.Atoi(pageNum)
	request.PagePerNums=int32(pageNumInt)

	keyword:=c.DefaultQuery("q","")
	request.KeyWords=keyword

	brandId:=c.DefaultQuery("b","0")
	brandIdInt,_:=strconv.Atoi(brandId)
	request.Brand=int32(brandIdInt)
	//**************ishot isnew istab
	isHot:=c.DefaultQuery("ih","0")
	if isHot=="1"{
		request.IsHot=true
	}
	isNew:=c.DefaultQuery("in","0")
	if isNew=="1"{
		request.IsNew=true
	}
	isTab:=c.DefaultQuery("it","0")
	if isTab=="1"{
		request.IsTab=true
	}

	zap.S().Info(global.GoodsSrvClient)
	rsp,err:=global.GoodsSrvClient.GoodsList(context.Background(),request)
	if err!=nil{
		zap.S().Errorw("[List] 查询【商品列表失败】")
		HandleGrpcErrorToHttp(err,c)
		return
	}
	goodsList:=make([]interface{},0)



	for _,value:=range rsp.Data{
		goodsList = append(goodsList, map[string]interface{}{
		"id": value.Id,
		"name":        value.Name,
		"goods_brief": value.GoodsBrief,
		"desc":        value.GoodsDesc,
		"ship_free":   value.ShipFree,
		"images":      value.Images,
		"desc_images": value.DescImages,
		"front_image": value.GoodsFrontImage,
		"shop_price":  value.ShopPrice,
		"ctegory": map[string]interface{}{
			"id":   value.Category.Id,
			"name": value.Category.Name,
		},
		"brand": map[string]interface{}{
			"id":   value.Brand.Id,
			"name": value.Brand.Name,
			"logo": value.Brand.Logo,
		},
		"is_hot":  value.IsHot,
		"is_new":  value.IsNew,
		"on_sale": value.OnSale,
	})
	}

	c.JSON(200,map[string]interface{}{
		"total":rsp.Total,
		"data":goodsList,
	})


}

