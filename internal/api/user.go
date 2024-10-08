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
//	@Description	可修改用户昵称
//	@Tags			User
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			nickname	formData	string	true	"昵称"
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

// BindEmail godoc
//
//	@Summary		绑定邮箱
//	@Description	发送邮件到用户指定邮箱，用户确认后进行绑定
//	@Tags			User
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			email	formData	string	true	"邮箱"
//	@Success		200		{object}	common.Response{data=response.BindEmailResp}
//	@Router			/user/email/bind [post]
func BindEmail(c *gin.Context) {
	var req *request.BindEmailReq
	if err := c.ShouldBind(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.UserService.BindEmail(c, req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// ValidEmail godoc
//
//	@Summary		邮箱绑定确认
//	@Description	通过指定链接确认绑定操作
//	@Tags			User
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Produce		json
//	@Param			token	query		string	true	"email token"
//	@Success		200		{object}	common.Response{data=response.ValidEmailResp}
//	@Router			/user/email/valid [get]
func ValidEmail(c *gin.Context) {
	var req request.ValidEmailReq
	if err := c.ShouldBind(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.UserService.ValidEmail(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// UserFollow godoc
//
//	@Summary		关注用户
//	@Description	基于id关注的其他用户
//	@Tags			User
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Accept			json
//	@Produce		json
//	@Param			id	body		string	true	"关注用户的id"
//	@Success		200	{object}	common.Response{data=response.UserFollowResp}
//	@Router			/user/follow [post]
func UserFollow(c *gin.Context) {
	var req request.UserFollowReq
	if err := c.ShouldBind(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.UserService.UserFollow(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// UserUnfollow godoc
//
//	@Summary		取关用户
//	@Description	基于id取关用户
//	@Tags			User
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Accept			json
//	@Produce		json
//	@Param			id	body		string	true	"取关用户的id"
//	@Success		200	{object}	common.Response{data=response.UserUnfollowResp}
//	@Router			/user/follow [delete]
func UserUnfollow(c *gin.Context) {
	var req request.UserUnfollowReq
	if err := c.ShouldBind(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.UserService.UserUnfollow(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// UserFollowingList godoc
//
//	@Summary		获取关注列表
//	@Description	获取用户的关注列表，列表中的用户信息包括头像和昵称
//	@Tags			User
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	common.Response{data=response.UserFollowingListResp}
//	@Router			/user/following [get]
func UserFollowingList(c *gin.Context) {
	var req request.UserFollowingListReq

	resp, code, isLogicError := service.UserService.UserFollowingList(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// UserFollowerList godoc
//
//	@Summary		获取粉丝列表
//	@Description	获取用户的粉丝列表，列表中的用户信息包括头像和昵称
//	@Tags			User
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	common.Response{data=response.UserFollowerListResp}
//	@Router			/user/follower [get]
func UserFollowerList(c *gin.Context) {
	var req request.UserFollowerListReq

	resp, code, isLogicError := service.UserService.UserFollowerList(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}
