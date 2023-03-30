package forms

type SendSmsForm struct{
	//mobile_validator 对应验证器 tag
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile_validator"`
	//1 登陆 2注册  用登陆 的验证码  去注册没有影响❓
	Type uint `form:"type" json:"type" binding:"required,oneof=1 2"`
}
