package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)
type Register struct{
	Host string
	Port int
}
type RegisterClient interface{
	Register(address string, port int, name string, tags []string, id string)(error)
	DeRegister(id string)error
}
func NewRegister(host string,port int)RegisterClient{
	return &Register{
		Host: host,
		Port: port,
	}

}
func (r *Register)Register(address string, port int, name string, tags []string, id string)(error) {
	fmt.Println(address,port,name)
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",r.Host,r.Port)

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
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	//registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}
func (r *Register)DeRegister(id string)error{
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",r.Host,r.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	err = client.Agent().ServiceDeregister(id)
	if err != nil {
		panic(err)
	}
	return nil
}