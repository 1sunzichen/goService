package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router:=gin.Default()
	router.GET("/welcome",welcome)
	router.POST("/postForm",postForm)
	router.Run(":8085")
}
func postForm(c *gin.Context){
	message:=c.PostForm("message")
	nick:=c.DefaultPostForm("nick","nick")
	c.JSON(http.StatusOK,gin.H{
		"message":message,
		"nick":nick,
	})
}
func welcome(c *gin.Context){
	firstname:=c.DefaultQuery("firstname","f")
	lastname:=c.DefaultQuery("lastname","l")
	c.JSON(200,gin.H{
		"lastname":lastname,
		"firstname":firstname,
	})
}