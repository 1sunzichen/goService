package forms

type PassWordLoginForm struct{
	//mobile_validator 对应验证器 tag
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile_validator"`
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=10"`
	Captcha string `form:"captcha" json:"captcha" binding:"required,min=3,max=10"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}
type RegisterForm struct{
	//mobile_validator 对应验证器 tag
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile_validator"`
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=10"`
	Code string `form:"code" json:"code" binding:"required"`
}
