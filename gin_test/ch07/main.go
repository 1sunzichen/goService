package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignUpForm struct{
	Age uint8 `json:"age" binding:"gte=1,lte=30"`
	Name string `json:"name" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"`
}
func main(){
	router:=gin.Default()
	router.POST("/signup", func(c *gin.Context) {
		var sign SignUpForm
		if err:=c.ShouldBind(&sign);err!=nil{
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"msg":"注册成功",
		})
	})
	router.Run(":8087")
}
