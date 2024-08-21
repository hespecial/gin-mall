package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/internal/api/request"
	"github.com/hespecial/gin-mall/internal/common"
	"github.com/hespecial/gin-mall/internal/common/e"
	"github.com/hespecial/gin-mall/internal/service"
	"github.com/hespecial/gin-mall/pkg/validator"
)

// ShowUserInfo godoc
//
//	@Summary		查看用户信息
//	@Description	可查看信息包括：用户名、昵称、用户状态、邮箱、头像
//	@Tags			User
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	common.Response{data=response.ShowUserInfoResp}
//	@Router			/user/info [get]
func ShowUserInfo(c *gin.Context) {
	var req *request.ShowUserInfoReq

	resp, code, isLogicError := service.UserService.ShowUserInfo(c, req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// UserInfoUpdate godoc
//
//	@Summary		修改用户信息
//	@Description	可修改昵称和邮箱
//	@Tags			User
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			nickname	formData	string	true	"昵称"
//	@Param			email		formData	string	true	"邮箱"
//	@Success		200			{object}	common.Response{data=response.UserInfoUpdateResp}
//	@Router			/user/info [put]
func UserInfoUpdate(c *gin.Context) {
	var req *request.UserInfoUpdateReq
	if err := c.ShouldBind(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.UserService.UserInfoUpdate(c, req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// UserPasswordChange godoc
//
//	@Summary		更改用户密码
//	@Description	输入`原密码-新密码-确认密码`以更改密码
//	@Tags			User
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			origin_password		formData	string	true	"原密码"
//	@Param			new_password		formData	string	true	"新密码"
//	@Param			confirm_password	formData	string	true	"确认密码"
//	@Success		200					{object}	common.Response{data=response.UserPasswordChangeResp}
//	@Router			/user/password [put]
func UserPasswordChange(c *gin.Context) {
	var req *request.UserPasswordChangeReq
	if err := c.ShouldBind(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.UserService.UserPasswordChange(c, req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// UploadAvatar godoc
//
//	@Summary		上传用户头像
//	@Description	上传头像，文件类型支持jpg(jpeg)、png
//	@Tags			User
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Accept			mpfd
//	@Produce		json
//	@Param			avatar	formData	file	true	"头像"
//	@Success		200		{object}	common.Response{data=response.UploadAvatarResp}
//	@Router			/user/avatar [post]
func UploadAvatar(c *gin.Context) {
	avatar, err := c.FormFile("avatar")
	if err != nil {
		common.Fail(c, e.ErrorUploadAvatar, e.IsLogicError)
		return
	}

	resp, code, isLogicError := service.UserService.UploadAvatar(c, avatar)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}
