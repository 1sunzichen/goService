package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"net"
	"os"
)
type Account struct{
	FirstName string `json:"firstname"`
}
func GETIPV4(){
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}

}
func main(){
	GETIPV4()
	host:="http://127.0.0.1:9200"
	 logger:=log.New(os.Stdout,"zc_shop",log.LstdFlags)
	client,err:=elastic.NewClient(elastic.SetURL(host),elastic.SetSniff(false),
		elastic.SetTraceLog(logger))
	if err!=nil{
		panic(err)
	}
	q:=elastic.NewMatchQuery("address","street")
	result,_:=client.Search().Index("user").Query(q).Do(context.Background())
	total:=result.Hits.TotalHits.Value
	fmt.Printf("%d",total)
	//for _,value:=range result.Hits.Hits{
	//	// json 数据结构转 struct
	//	var acc Account
	//	_=json.Unmarshal(value.Source,&acc)
	//	fmt.Println(acc.FirstName,"name")
	//	jsonData,_:=value.Source.MarshalJSON();
	//		fmt.Println(string(jsonData))
	//
	//}
	account:=Account{FirstName: "小明"}
	put1,_:=client.Index().Index("myuser").BodyJson(account).Do(context.Background())
	fmt.Println(put1.Id)
}

