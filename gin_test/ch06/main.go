package main

import (
	"github.com/gin-gonic/gin"
	proto "gopro/gin_test/ch06/protobuf/api"
)

func main(){
	router:=gin.Default()
	router.POST("/protobufPost",returnProto)
	router.Run(":8086")
}

func returnProto(c *gin.Context){
	//course:=[]string{"123","111"}
	course:=c.PostFormArray("Course")

	user:=&proto.Teacher{
		Name:"sjq",
		Course:course,
	}
	c.ProtoBuf(200,user)
}