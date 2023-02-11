package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignUpForm struct{
	Age uint8 `json:"age" binding:"gte=1,lte=130"`
	Name string `json:"name" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

func removeTopStruct(){

}

func main(){
   router:=gin.Default()
   router.POST("/signup", func(c *gin.Context) {
	  var signUpForm  SignUpForm
	  if err:=c.ShouldBind(&signUpForm);err!=nil{
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
}
