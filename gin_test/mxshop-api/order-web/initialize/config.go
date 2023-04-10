package initialize

import (
	"fmt"
	viper "github.com/spf13/viper"

	"github.com/fsnotify/fsnotify"
	"gopro/gin_test/mxshop-api/order-web/global"

	//"time"
)
type MysqlConfig struct{
	host string `mapstructure:"host"`
	port string `mapstructure:"port"`
}
type ServerStruct struct{
	ServerName string `mapstructure:"name"`
}
func GetEnvInfo(env string)bool{
	viper.AutomaticEnv()
	return viper.GetBool(env)
}
func InitConfig(){
	debug:=GetEnvInfo("MXSHOP_DEBUG")
	configPrefix:="config"
	configName:=fmt.Sprintf("gin_test/mxshop-api/order-web/%spro.yaml",configPrefix)
	v:=viper.New()
	//文件路径如何设置
	fmt.Println(debug,"debug")
	if debug{
		configName=fmt.Sprintf("gin_test/mxshop-api/order-web/%sbug.yaml",configPrefix)
	}

	v.SetConfigFile(configName)

	if err:=v.ReadInConfig();err!=nil{
		panic(err)
	}
	if err:=v.Unmarshal(&global.ServerConfig);err!=nil{
		panic(err)
	}
	fmt.Println(v.Get("name"),GetEnvInfo("MXSHOP_DEBUG"))
	//
	v.WatchConfig()
	//fsnotify.Event
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config",in.Name)
		_=v.ReadInConfig()
		_=v.Unmarshal(&global.ServerConfig)
		fmt.Println(global.ServerConfig)
	})
	//time.Sleep(time.Second*3000)
}

