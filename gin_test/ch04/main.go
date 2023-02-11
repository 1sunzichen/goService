package main

import (
	"github.com/gin-gonic/gin"
)
type Person struct{
	ID string `uri:"id" binding:"required,uuid"`
}
func main(){
	//uuid 13ebbb93-0625-4daf-813b-4a0e81461f6d
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/:id", loginEndpoint)
	}
	router.GET("/:id", loginEndpoint)
	// Simple group: v2

	router.Run(":8082")
}
func loginEndpoint(c *gin.Context){
	var person Person
	if err:=c.ShouldBindUri(&person);err!=nil{
		c.Status(404)
		return
	}

	//id:=c.Param("id")
	c.JSON(200,gin.H{
		"id":person.ID,
	})
}