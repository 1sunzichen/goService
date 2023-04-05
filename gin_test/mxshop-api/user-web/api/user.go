package api

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopro/gin_test/mxshop-api/user-web/forms"
	"gopro/gin_test/mxshop-api/user-web/global"
	"gopro/gin_test/mxshop-api/user-web/global/response"
	"gopro/gin_test/mxshop-api/user-web/middlewares"
	"gopro/gin_test/mxshop-api/user-web/models"
	"gopro/gin_test/mxshop-api/user-web/proto"

	//"gopro/gin_test/mxshop-api/user-web/proto"

	"net/http"
	"strconv"
	"strings"
	"time"
)
func removeTopStruct(fields map[string]string) map[string]string{
	rsp:=map[string]string{}
	for field,err:=range fields{
		rsp[field[strings.Index(field,".")+1:]]=err
	}
	return rsp
}
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

func HandleVaildatorError(c *gin.Context,err error){
	errs,ok:=err.(validator.ValidationErrors)
	if !ok{
		c.JSON(http.StatusOK,gin.H{
			"msg":err.Error(),
		})
	}
	fmt.Println(errs,"错误")
	c.JSON(http.StatusBadRequest,gin.H{
		"error":removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}
func GetConsul()(string,int){
	var address string
	var port int
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",global.ServerConfig.ConsulInfo.Host,8500)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
    un:=global.ServerConfig.UserSrvInfo.Name
    fmt.Printf("un%s",un)
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"",un))
	//_, err = client.Agent().ServicesWithFilter(`Service == "user_srv"`)
	if err != nil {
		panic(err)
	}
	//
 	for key, value := range data {
		fmt.Println(key,value,value.Address,value.Port)
		address=value.Address
		port=value.Port
		//break
	}
	return  address,port
}
func GetUserList(c *gin.Context){
	//注册中心
	//host,port:= GetConsul()
	//ip:="127.0.0.1"
	//port:=50051
	//拨号链接用户服务器
	//userConn,err:=grpc.Dial(fmt.Sprintf("%s:%d",global.ServerConfig.UserSrvInfo.Host,global.ServerConfig.UserSrvInfo.Port ),grpc.WithInsecure())

	claims,_:=c.Get("claims")
	currentUser:=claims.(*models.CustomClaims)
	zap.S().Infof("访问用户：%d",currentUser.Id)
	//生成grpc的client并调用接口

	//
	pn:=c.DefaultQuery("pn","0")
	//string 转 int
	pnInt,_:=strconv.Atoi(pn)
	pSize:=c.DefaultQuery("psize","10")
	pSizeInt,_:=strconv.Atoi(pSize)

	rsp,err:=global.UserSrvClient.GetUserList(context.Background(),&proto.PageInfo{
		Pn: uint32(pnInt),
		PSize: uint32(pSizeInt),
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

func PasswordLogin(c *gin.Context){
	passwordLoginForm:=forms.PassWordLoginForm{}

	if err:=c.ShouldBind(&passwordLoginForm);err!=nil{
		HandleVaildatorError(c,err)
		return
	}
	if !store.Verify(passwordLoginForm.CaptchaId,passwordLoginForm.Captcha,true){
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"验证码错误",
		})
		return
	}
	userConn,err:=grpc.Dial(fmt.Sprintf("%s:%d",global.ServerConfig.UserSrvInfo.Host,global.ServerConfig.UserSrvInfo.Port ),grpc.WithInsecure())
	if err!=nil{
		zap.S().Errorw("[getuserlist] 链接【用户服务失败】",
			"msg",err.Error())
	}
	fmt.Println(global.ServerConfig.UserSrvInfo.Host,global.ServerConfig.UserSrvInfo.Port )
	//生成grpc的client并调用接口
	userSrvClient:=proto.NewUserClient(userConn)
	if rsp,err:=userSrvClient.GetUserMobile(context.Background(),&proto.MobileReq{
		Mobile: passwordLoginForm.Mobile,
	});err!=nil{
		if e,ok:=status.FromError(err); ok {
			fmt.Println(e.Code(),"错误玛")
		switch e.Code() {
		case codes.NotFound:
			c.JSON(http.StatusBadRequest,map[string]string{
				"mobile":"用户不存在",
		   		})
		default:
			c.JSON(http.StatusInternalServerError,map[string]string{
			    "mobile":"登陆失败",
			})
	     }
	     return
	  }
	}else{
		//查询了用户，没有检查密码 校验与数据库的密码是否一致 rsp.PassWord 数据库的密码
		if passRsp,passErr:=userSrvClient.CheckPassWord(context.Background(),&proto.PassWordInfo{
			PassWord: passwordLoginForm.PassWord,
			EncryptedPassword: rsp.PassWord,
		});passErr!=nil{
			c.JSON(http.StatusInternalServerError,map[string]string{
				"password":"两次密码输入不一致",
			})
		}else{
			if passRsp.Success{
				//生成token
				jwtHandler:=middlewares.NewJWT()
				claims:=models.CustomClaims{
					ID:uint(rsp.Id),
					NickName: rsp.Nickname,
					AuthorityId: uint(rsp.Role),
					StandardClaims:jwt.StandardClaims{
						NotBefore: time.Now().Unix(),
						ExpiresAt: time.Now().Unix()+60*60*24*100,
						Issuer: "zc",
					},
				}
				token,err:=jwtHandler.CreateToken(claims)
				if err!=nil{

					c.JSON(http.StatusInternalServerError,map[string]string{
						"msg":"生成token失败",
					})
					return
				}

				c.JSON(http.StatusOK,map[string]interface{}{
					"msg":"登陆成功",
					"id":rsp.Id,
					"nick_name":rsp.Nickname,
					"token": token,
					"expired_at":time.Now().Unix()+60*60*24*1000*100,
				})
			}else{
				c.JSON(http.StatusBadRequest,map[string]string{
					"msg":"登陆失败-密码，网络问题",
				})
			}
		}

	}
}
func Register(c *gin.Context){
	//用户注册
	registerForm := forms.RegisterForm{}
	if err := c.ShouldBind(&registerForm); err != nil {
		HandleVaildatorError(c, err)
		return
	}

	//验证码
	rdb := redis.NewClient(&redis.Options{
		Addr:fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	value, err := rdb.Get(registerForm.Mobile).Result()
	if err == redis.Nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"code":"验证码错误",
		})
		return
	}else{
		if value != registerForm.Code {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":"验证码错误",
			})
			return
		}
	}
	//userConn,err:=grpc.Dial(fmt.Sprintf("%s:%d",global.ServerConfig.UserSrvInfo.Host,global.ServerConfig.UserSrvInfo.Port ),grpc.WithInsecure())
	//if err!=nil{
	//	zap.S().Errorw("[getuserlist] 链接【用户服务失败】",
	//		"msg",err.Error())
	//}
	//fmt.Println(global.ServerConfig.UserSrvInfo.Host,global.ServerConfig.UserSrvInfo.Port )
	////生成grpc的client并调用接口
	//userSrvClient:=proto.NewUserClient(userConn)
	user, err := global.UserSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		NickName: registerForm.Mobile,
		PassWord: registerForm.PassWord,
		Mobile:   registerForm.Mobile,
	})

	if err != nil {
		zap.S().Errorf("[Register] 查询 【新建用户失败】失败: %s", err.Error())
		HandleGrpcErrorToHttp(err, c)
		return
	}
    // 返回token信息
	j := middlewares.NewJWT()
	claims := models.CustomClaims{
		ID:             uint(user.Id),
		NickName:       user.Nickname,
		AuthorityId:    uint(user.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(), //签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
			Issuer: "imooc",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":"生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": user.Id,
		"nick_name": user.Nickname,
		"token": token,
		"expired_at": (time.Now().Unix() + 60*60*24*30)*1000,
	})
}