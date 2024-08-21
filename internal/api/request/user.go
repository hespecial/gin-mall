package request

import "github.com/hespecial/gin-mall/pkg/validator"

type UserInfoUpdateReq struct {
	Nickname string `form:"nickname" json:"nickname" binding:"required,max=32"`
}

func (*UserInfoUpdateReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		"Nickname.required": "昵称不能为空",
		"Nickname.max":      "昵称最大长度为32个字符",
	}
}

type UserPasswordChangeReq struct {
	OriginPassword  string `form:"origin_password" json:"origin_password" binding:"required"`
	NewPassword     string `form:"new_password" json:"new_password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required,eqfield=NewPassword"`
}

func (*UserPasswordChangeReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		"OriginPassword.required":  "原密码不能为空",
		"NewPassword.required":     "新密码不能为空",
		"ConfirmPassword.required": "确认密码不能为空",
		"ConfirmPassword.eqfield":  "确认密码不一致",
	}
}

type ShowUserInfoReq struct{}

type UpdateAvatarReq struct{}

type BindEmailReq struct {
	Email string `form:"email" json:"email" binding:"email"`
}

func (*BindEmailReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		"Email.email": "邮箱格式错误",
	}
}

type UnbindEmailReq struct{}

type ValidEmailReq struct {
	Token string `form:"token" binding:"required"`
}

func (*ValidEmailReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		"Token.required": "缺失token",
	}
}
