package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/internal/api/request"
	"github.com/hespecial/gin-mall/internal/common"
	"github.com/hespecial/gin-mall/internal/common/e"
	"github.com/hespecial/gin-mall/internal/service"
	"github.com/hespecial/gin-mall/pkg/validator"
)

// Register godoc
//
//	@Summary		用户注册
//	@Description	输入`用户名-密码-确认密码`以注册
//	@Tags			Auth
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			username			formData	string	true	"用户名"
//	@Param			password			formData	string	true	"密码"
//	@Param			confirm_password	formData	string	true	"确认密码"
//	@Success		200					{object}	common.Response{data=response.AuthRegisterResp}
//	@Router			/auth/register [post]
func Register(c *gin.Context) {
	var req request.AuthRegisterReq
	if err := c.ShouldBind(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.AuthService.Register(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// Login godoc
//
//	@Summary		用户登录
//	@Description	输入`用户名-密码`以登录
//	@Tags			Auth
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			username	formData	string	true	"用户名"
//	@Param			password	formData	string	true	"密码"
//	@Success		200			{object}	common.Response{data=response.AuthLoginResp}
//	@Router			/auth/login [post]
func Login(c *gin.Context) {
	var req request.AuthLoginReq
	if err := c.ShouldBind(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.AuthService.Login(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}
