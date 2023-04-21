package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main(){
	p,err:=rocketmq.NewProducer(producer.WithNameServer([]string{"192.168.155.14:9876"}))
	if err!=nil{
		panic("生成producer 失败")
	}
	 if err=p.Start();err!=nil{
		 panic(err.Error())
	 }

    res,err:=p.SendSync(context.Background(),primitive.NewMessage("imooc",[]byte("this is ")))
    if err!=nil{
    	panic("发送失败")
	}else{
		fmt.Printf("发送成功:%s\n",res.String())
	}

	if err=p.Shutdown();err!=nil{panic("关闭producer失败 ")}
}

