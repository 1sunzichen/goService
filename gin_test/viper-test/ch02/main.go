package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	viper "github.com/spf13/viper"
	"gopro/gin_test/mxshop-api/user-web/config"
	"time"
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
func main(){
	debug:=GetEnvInfo("MXSHOP_DEBUG")
	configPrefix:="config"
	configName:=fmt.Sprintf("gin_test/mxshop-api/user-web/%spro.yaml",configPrefix)
	v:=viper.New()
	//文件路径如何设置
	if debug{
		configName=fmt.Sprintf("gin_test/mxshop-api/user-web/%sbug.yaml",configPrefix)
	}

	v.SetConfigFile(configName)
	severStruct:=config.ServerConfig{}
	if err:=v.ReadInConfig();err!=nil{
		panic(err)
	}

	fmt.Println(v.Get("name"),GetEnvInfo("MXSHOP_DEBUG"))
	//
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config",in.Name)
		v.ReadInConfig()
		v.Unmarshal(&severStruct)
		fmt.Println(severStruct)
	})
	time.Sleep(time.Second*3000)
}
