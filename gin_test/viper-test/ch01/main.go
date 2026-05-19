package main

import (
	"fmt"
	"github.com/spf13/viper"
)
type MysqlConfig struct{
	host string `mapstructure:"host"`
	port string `mapstructure:"port"`
}
type ServerStruct struct{
	ServerName string `mapstructure:"name"`
}
func main(){
	v:=viper.New()
	// configure file path
	v.SetConfigFile("configbug.yaml")
	if err:=v.ReadInConfig();err!=nil{
		panic(err)
	}
	severStruct:=ServerStruct{}
	if err:=v.Unmarshal(&severStruct);err!=nil{
		panic(err)
	}
	fmt.Println(v.Get("name"),severStruct)
}
