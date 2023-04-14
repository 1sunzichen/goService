package main
import (
	"fmt"
	//"github.com/nacos-group/nacos-sdk-go/v2"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)
func main(){
	serverConfigs:=[]constant.ServerConfig{
		{
			IpAddr:"192.168.3.21",
			Port: 8848,
		},
	}
	clientConfig := constant.ClientConfig{
		NamespaceId:         "9529f44a-dd4a-4242-817e-d35f3a726128", // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./",
		CacheDir:            "./",
		LogLevel:            "debug",
	}
	// 创建动态配置客户端
	configClient, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user-web.yaml",
		Group:  "DEFAULT_GROUP"})
	if err!=nil{
		panic(err)
	}
	fmt.Println(content)
}
