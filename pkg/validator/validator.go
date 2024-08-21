package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type ValidateMessages map[string]string

type Validator interface {
	ErrorMessages() ValidateMessages
}

// GetErrorMsg 获取错误信息
func GetErrorMsg(request interface{}, err error) (msg string) {
	msg = "invalid parameter"
	if !errors.As(err, &validator.ValidationErrors{}) {
		return
	}

	v, isValidator := request.(Validator)
	if !isValidator {
		return
	}
	// 若 request 实现 Validator 接口即可实现自定义错误信息
	// 返回第一个字段的错误
	field := err.(validator.ValidationErrors)[0]
	if message, exist := v.ErrorMessages()[field.Field()+"."+field.Tag()]; exist {
		return message
	}
	// 无法映射到 ValidatorMessages 中的元素时返回原错误
	return
}
