package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopro/gin_test/mxshop-api/order-web/global"
	"gopro/gin_test/mxshop-api/order-web/initialize"
	"gopro/gin_test/mxshop-api/order-web/utils"
	"gopro/gin_test/mxshop-api/order-web/utils/register/consul"
	validatorss "gopro/gin_test/mxshop-api/order-web/validator"
	"os"
	"os/signal"
	"syscall"
)
//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6NTksIk5pY2tOYW1lIjoiYm9kZHk5IiwiQXV0aG9yaXR5SWQiOjEsImV4cCI6MTY4ODQzNTA1OSwiaXNzIjoiemMiLCJuYmYiOjE2Nzk3OTUwNTl9.-OwFSUC9AOr0OMDDt6D8pkqan2YSwzUg6bgCZJd-Adc
func main(){
	port:=8023
	//1.初始化logger
	initialize.InitLogger()
	//2.初始化配置文件
	initialize.InitConfig()

	//3.初始化routers
	r:=initialize.Routers()

	//4.初始化翻译
	if err:=initialize.InitTrans("zh");err!=nil{
	 	panic(err)
	 }
	//5 初始化 获取注册中心src链接
	//initialize.InitConsul()
	//5 负载均衡
	initialize.InitSrvConn()
	a,err:=utils.GetFreePort()
	if err==nil{
		fmt.Println(a,"动态port")
	}
	viper.AutomaticEnv()
	debug:=viper.GetBool("MXSHOP_DEBUG")
	if !debug{
		port,err:=utils.GetFreePort()
		if err==nil{
			global.ServerConfig.Port=port
		}
	}
	//
	 //注册验证器
	 if v,ok:=binding.Validator.Engine().(*validator.Validate);ok{
	 	_=v.RegisterValidation(global.MobileValidator,validatorss.ValidateMobile)
	 	_=v.RegisterTranslation(global.MobileValidator, global.Trans, func(ut ut.Translator) error {
			 return ut.Add(global.MobileValidator, "{0}格式不正确!", true) // see universal-translator for details
		 }, func(ut ut.Translator, fe validator.FieldError) string {
			 t, _ := ut.T(global.MobileValidator, fe.Field())

			 return t
		 })
	 }

	 registerClient :=consul.NewRegister(global.ServerConfig.ConsulInfo.Host,global.ServerConfig.ConsulInfo.Port)
	 serviceId:=fmt.Sprintf("%s",uuid.NewV4())
	 err=registerClient.Register("127.0.0.1",global.ServerConfig.Port,global.ServerConfig.Name,[]string{"web服务","订单服务"},serviceId)
	 if err!=nil{
	 	zap.S().Panic("服务注册失败",err.Error())
	 }
	 go func() {
		 if err:=r.Run(fmt.Sprintf(":%d",port));err!=nil{
			 zap.S().Panic("启动失败",err.Error())
		 }
	 }()
	quit:=make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	servicesId:=[]string{
		//"00c7325a-9818-48a8-9935-c30ecd7b12af",
	}
	for _,s:=range servicesId {
		if err = registerClient.DeRegister(s); err != nil {
			zap.S().Info("服务发现注销失败")
		} else {
			zap.S().Info("服务发现注销成功")
		}
	}
	if err=registerClient.DeRegister(serviceId);err!=nil{
		zap.S().Info("服务发现注销order-web失败")
	}else{
		zap.S().Info("服务发现注销order-web成功")
	}

}
