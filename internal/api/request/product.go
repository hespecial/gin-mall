package request

import "github.com/hespecial/gin-mall/pkg/validator"

type GetProductListReq struct {
	Page int `form:"page" json:"page" binding:"required,min=1"`
	Size int `form:"size" json:"size" binding:"required,min=1"`
}

func (*GetProductListReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		"Page.required": "页号不能为空",
		"Page.min":      "页号最小为1",
		"Size.max":      "页码不能为空",
		"Size.min":      "页码最小为1",
	}
}

type GetProductDetailInfoReq struct {
	ID uint `uri:"id" binding:"required,min=1"`
}

func (*GetProductDetailInfoReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		"ID.required": "参数有误",
		"ID.min":      "参数有误",
	}
}

type SearchProductReq struct {
	Keyword string `form:"keyword" json:"keyword" binding:"required"`
	Page    int    `form:"page" json:"page" binding:"required,min=1"`
	Size    int    `form:"size" json:"size" binding:"required,min=1"`
}

func (*SearchProductReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{}
}
