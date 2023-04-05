package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopro/gin_test/mxshop-api/goods-web/models"
)

func IsAdminAuth() gin.HandlerFunc{
	return func(context *gin.Context) {
		claims,_:=context.Get("claims")
		currentUser:=claims.(*models.CustomClaims)
		fmt.Println(currentUser.AuthorityId,"currentUser.AuthorityId")
		//if currentUser.AuthorityId!=2{
		//	context.JSON(http.StatusForbidden,gin.H{
		//		"msg":"无权限",
		//	})
		//	context.Abort()
		//	return
		//}
		context.Next()
	}
}