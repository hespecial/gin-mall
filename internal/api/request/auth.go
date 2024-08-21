package request

import "github.com/hespecial/gin-mall/pkg/validator"

type AuthRegisterReq struct {
	Username        string `form:"username" json:"username" binding:"required,max=32"`
	Password        string `form:"password" json:"password" binding:"required,min=6,max=20"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required,eqfield=Password"`
}

func (*AuthRegisterReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		"Username.required":        "用户名不能为空",
		"Username.max":             "用户名长度不能超过32个字符",
		"Password.required":        "密码不能为空",
		"Password.min":             "密码长度范围在6-20个字符",
		"Password.max":             "密码长度范围在6-20个字符",
		"ConfirmPassword.required": "确认密码不能为空",
		"ConfirmPassword.eqfield":  "确认密码不一致",
	}
}

type AuthLoginReq struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (*AuthLoginReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		"Username.required": "用户名不能为空",
		"Password.required": "密码不能为空",
	}
}
