package common

import (
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/internal/common/e"
	"net/http"
)

type Response struct {
	Code e.Code      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	const codeSuccess = e.Success
	resp := &Response{
		Code: codeSuccess,
		Msg:  codeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, resp)
}

func FailWithMsg(c *gin.Context, code e.Code, msg string) {
	resp := &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, resp)
}

func Fail(c *gin.Context, code e.Code, isLogicError bool) {
	if !isLogicError {
		FailWithMsg(c, code, code.Msg())
	} else {
		FailWithMsg(c, code, e.BusinessLogicError)
	}
}
