package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	viper "github.com/spf13/viper"
	"gopro/gin_test/inventory_srv/global"

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
	configName:=fmt.Sprintf("gin_test/inventory_srv/%spro.yaml",configPrefix)
	v:=viper.New()
	//文件路径如何设置
	if debug{
		configName=fmt.Sprintf("gin_test/inventory_srv/%sbug.yaml",configPrefix)
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
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config",in.Name)
		_=v.ReadInConfig()
		_=v.Unmarshal(&global.ServerConfig)
		fmt.Println(&global.ServerConfig)
	})
	//time.Sleep(time.Second*3000)
}

