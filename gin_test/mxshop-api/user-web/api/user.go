package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopro/gin_test/mxshop-api/user-web/global"
	"gopro/gin_test/mxshop-api/user-web/global/response"
	"gopro/gin_test/mxshop_srv/proto"
	"net/http"
	"strconv"
	"time"
)

func HandleGrpcErrorToHttp(err error,c *gin.Context){
	//将grpc 的code转换成http的状态玛
	if err!=nil{
		if e,ok:=status.FromError(err);ok{
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})

			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": "参数",
				})
			default:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": "其他",
				})
			}
			return

		}
	}

}

func GetUserList(c *gin.Context){
	//ip:="127.0.0.1"
	//port:=50051
	//拨号链接用户服务器
	userConn,err:=grpc.Dial(fmt.Sprintf("%s:%d",global.ServerConfig.UserSrvInfo.Host,global.ServerConfig.UserSrvInfo.Port ),grpc.WithInsecure())
	if err!=nil{
		zap.S().Errorw("[getuserlist] 链接【用户服务失败】",
			"msg",err.Error())
	}
	//生成grpc的client并调用接口
	userSrvClient:=proto.NewUserClient(userConn)
	//
	pn:=c.DefaultQuery("pn","0")
	pnInt,_:=strconv.Atoi(pn)
	pSize:=c.DefaultQuery("psize","10")
	pSizeInt,_:=strconv.Atoi(pSize)

	rsp,err:=userSrvClient.GetUserList(context.Background(),&proto.PageInfo{
		Pn: 1,
		PSize: 2,
	})
	if err!=nil{
		zap.S().Errorw("[GetUserList] 查询【用户列表失败】")
		HandleGrpcErrorToHttp(err,c)
		return
	}
	result:=make([]interface{},0)

	for _,value:=range rsp.Data{
		user:=response.UserResponse{
			Id:value.Id,
			NickName: value.Nickname,
			//BirthDay: time.Time(time.Unix(int64(value.Birthday),0)).Format("2006-01-02"),
			BirthDay: response.JsonTime(time.Unix(int64(value.Birthday),0)),
			//BirthDay: time.Time(time.Unix(int64(value.Birthday),0)),
			Gender: value.Gender,
			Mobile: value.Mobile,

		}
		//data:=make(map[string]interface{})
		//data["id"]=value.Id
		//data["name"]=value.Nickname
		//data["birthday"]=value.Birthday
		//data["gender"]=value.Gender
		result=append(result,user)
	}
	c.JSON(http.StatusOK,result)
	zap.S().Debug("获取用户列表页")

}