package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/satori/go.uuid"
	"gopro/gin_test/mxshop_srv/global"
)

func InitConsul(port int)(*api.Client,string){
	serverId:=fmt.Sprintf("%s",uuid.NewV4())
	client,err:=Register(global.ServerConfig.ConsulInfo.Host,global.ServerConfig.ConsulInfo.Port,global.ServerConfig.Name,[]string{"123","123"},serverId,port)
	if err!=nil{
		return nil,""
	}
	return client,serverId
}
func Register(address string, port int, name string, tags []string, id string,srcPort int)(*api.Client, error) {
	fmt.Println(address,port,name)
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",address,port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	//check := &api.AgentServiceCheck{
	//	GRPC:                           "http://192.168.3.23:50051",
	//	Timeout:                        "5s",
	//	Interval:                       "5s",
	//	DeregisterCriticalServiceAfter: "10s",
	//}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	//registration.Port = 50051
	registration.Port = srcPort
	registration.Tags = tags
	registration.Address = address
	//registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return client,nil
}