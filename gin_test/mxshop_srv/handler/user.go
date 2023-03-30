package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopro/gin_test/mxshop_srv/global"
	"gopro/gin_test/mxshop_srv/model"
	"gopro/gin_test/mxshop_srv/proto"
	"gorm.io/gorm"
	"strings"
	"time"
)


type UserServer struct{
	proto.UnimplementedUserServer
}



func ModelToRep(user model.User)proto.UserInfoRes{
	//grpc 的message字段 有默认值 你不能随便赋值nil进去，容易出错
	// 这里要搞清，哪些字段是有默认值的
	userInfoResponse:= proto.UserInfoRes{
		Id: user.ID,
		PassWord: user.Password,
		Nickname: user.NickName,
		Gender:user.Gender,
		Role:int32(user.Role),
	}
	if user.Birthday!=nil{
		userInfoResponse.Birthday=uint64(user.Birthday.Unix())
	}
	return userInfoResponse
}
func (c *UserServer) GetUserList(ctx context.Context,req *proto.PageInfo)(*proto.UserListRes,error){
	var users []model.User
	result :=global.DB.Find(&users)
	if result.Error!=nil{
		return nil,result.Error
	}
	//
	fmt.Println("用户列表******")
	rsp:=&proto.UserListRes{

	}
	rsp.Total=int32(result.RowsAffected)
	global.DB.Scopes(Paginate(int(req.Pn),int(req.PSize))).Find(&users)
    for _,user:=range  users{
    	userInfoRsp:=ModelToRep(user)
    	rsp.Data=append(rsp.Data,&userInfoRsp)
	}
	return rsp,nil
}
func Paginate(page ,pageSize int) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		//q := r.URL.Query()
		//page, _ := strconv.Atoi(q.Get("page"))
		if page == 0 {
			page = 1
		}
		//pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}

}


func (c *UserServer) GetUserMobile(ctx context.Context, req *proto.MobileReq) (*proto.UserInfoRes, error){
	var user model.User
	result:=global.DB.Where(&model.User{Mobile:req.Mobile }).First(&user)
	if result.RowsAffected==0{
		return nil,status.Errorf(codes.NotFound,"用户不存在")
	}
	if result.Error!=nil{
		return nil,result.Error
	}
	userRep:=ModelToRep(user)
	return &userRep,nil
}
func (c *UserServer) GetUserId(ctx context.Context, req *proto.IdReq) (*proto.UserInfoRes, error){
	//通过ID 查询用户
	var user model.User
	result:=global.DB.First(&user,req.Id)
	if result.RowsAffected==0{
		return nil,status.Errorf(codes.NotFound,"用户不存在")
	}
	if result.Error!=nil{
		return nil,result.Error
	}
	userRep:=ModelToRep(user)
	return &userRep,nil
}

func (c *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoRes, error){
	var user model.User
	result:=global.DB.Where(&model.User{Mobile:req.Mobile }).First(&user)
	if result.RowsAffected==1{
		return nil,status.Errorf(codes.AlreadyExists,"用户已存在")
	}
	user.Mobile=req.Mobile
	user.NickName=req.NickName
	//
	// Using custom options
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(req.PassWord, options)
	user.Password=fmt.Sprintf("$pbkdf2-sha512$%s$%s",salt,encodedPwd)
	result=global.DB.Create(&user)
	if result.Error!=nil{
		return nil,status.Errorf(codes.Internal,result.Error.Error())
	}
	userRep:=ModelToRep(user)
	return &userRep,nil
}
func (c *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*empty.Empty, error){
	//
	var user model.User
	result:=global.DB.First(&user,req.Id)
	if result.RowsAffected==0{
		return nil,status.Errorf(codes.NotFound,"用户不存在")
	}
	birthday:=time.Unix(int64(req.Birthday),0)
	user.Birthday=&birthday
	user.NickName=req.Nickname
	user.Gender=req.Gender
	result=global.DB.Save(user)
	if result.Error!=nil{
		return nil,status.Errorf(codes.Internal,result.Error.Error())
	}
	return &empty.Empty{},nil
}
func (c *UserServer) CheckPassWord(ctx context.Context, req *proto.PassWordInfo) (*proto.CheckRes, error) {
	//校验密码
	passwordInfo:=strings.Split(req.EncryptedPassword,"$")
	options := &password.Options{16, 100, 32, sha512.New}

	check:=password.Verify(req.PassWord,passwordInfo[2],passwordInfo[3],options)
	return &proto.CheckRes{Success: check},nil

}