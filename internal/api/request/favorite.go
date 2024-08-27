package request

import "github.com/hespecial/gin-mall/pkg/validator"

type GetFavoriteListReq struct{}

type AddFavoriteReq struct {
	ID uint `json:"id" binding:"required"`
}

func (*AddFavoriteReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		"ID.required": "参数有误",
	}
}

type DeleteFavoriteReq struct {
	ID uint `json:"id" binding:"required"`
}

func (*DeleteFavoriteReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		"ID.required": "参数有误",
	}
}
