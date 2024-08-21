package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/api/request"
	"github.com/hespecial/gin-mall/internal/api/response"
	"github.com/hespecial/gin-mall/internal/common/e"
	"github.com/hespecial/gin-mall/internal/common/enum"
	"github.com/hespecial/gin-mall/internal/model"
	"github.com/hespecial/gin-mall/internal/repository/dao"
	"github.com/hespecial/gin-mall/pkg/files"
	"github.com/hespecial/gin-mall/pkg/oss"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mime/multipart"
	"path"
	"strings"
)

type userService struct{}

var UserService *userService

func getUserID(c *gin.Context) (uint, bool) {
	userID, exist := c.Get(enum.UserID)
	if !exist {
		global.Log.Error("jwt中间件上下文值传递错误")
		return 0, false
	}
	return userID.(uint), true
}

func getUsername(c *gin.Context) (string, bool) {
	username, exist := c.Get(enum.Username)
	if !exist {
		global.Log.Error("jwt中间件上下文值传递错误")
		return "", false
	}
	return username.(string), true
}

func getAvatarURL(updateMode string, filename string) (url string) {
	switch updateMode {
	case enum.UploadModeLocal:
		url = fmt.Sprintf("http://%s:%d/%s",
			global.Config.Server.Host,
			global.Config.Server.Port,
			global.Config.Image.AvatarDir,
		)
	case enum.UploadModeOSS:
		url = fmt.Sprintf("https://%s.%s",
			global.Config.Oss.Bucket,
			global.Config.Oss.Endpoint,
		)
	}
	return strings.Join([]string{url, filename}, "/")
}

func (*userService) ShowUserInfo(c *gin.Context, _ *request.ShowUserInfoReq) (*response.ShowUserInfoResp, e.Code, bool) {
	// 从context中获取UserID
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	// 根据UserID获取用户
	user, err := dao.GetUserByID(userID)
	if err != nil {
		global.Log.Error("根据ID获取用户失败", zap.Error(err))
		return nil, e.ErrorGetUserByID, e.IsLogicError
	}

	// 响应数据
	resp := &response.ShowUserInfoResp{
		Username: user.Username,
		Nickname: user.Nickname,
		Status:   user.Status,
		Email:    user.Email,
		Avatar:   user.Avatar,
	}

	return resp, e.Success, e.NotLogicError
}

func (*userService) UserInfoUpdate(c *gin.Context, req *request.UserInfoUpdateReq) (*response.UserInfoUpdateResp, e.Code, bool) {
	// 从context中获取UserID
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	// 创建更新实例
	userForUpdate := &model.User{
		Model:    gorm.Model{ID: userID},
		Nickname: req.Nickname,
		Email:    req.Email,
	}

	// 更新用户信息
	err := dao.UpdateUser(userForUpdate)
	if err != nil {
		global.Log.Error("用户信息更新失败", zap.Error(err))
		return nil, e.ErrorUpdateUser, e.IsLogicError
	}

	// 响应信息
	resp := &response.UserInfoUpdateResp{}

	return resp, e.Success, e.NotLogicError
}

func (*userService) UserPasswordChange(c *gin.Context, req *request.UserPasswordChangeReq) (*response.UserPasswordChangeResp, e.Code, bool) {
	// 从context中获取UserID
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	// 根据UserID获取用户
	user, err := dao.GetUserByID(userID)
	if err != nil {
		global.Log.Error("根据ID获取用户失败", zap.Error(err))
		return nil, e.ErrorGetUserByID, e.IsLogicError
	}

	// 密码校验
	valid := user.CheckPassword(req.OriginPassword)
	if !valid {
		// 原密码错误
		global.Log.Info("用户更改密码失败")
		return nil, e.ErrorIncorrectPassword, e.NotLogicError
	}

	// 创建更新实例
	userForUpdate := &model.User{
		Model:    gorm.Model{ID: userID},
		Password: req.NewPassword,
	}

	// 密码加密
	err = userForUpdate.EncryptPassword()
	if err != nil {
		global.Log.Error("密码加密失败", zap.Error(err))
		return nil, e.ErrorEncryptPassword, e.IsLogicError
	}

	// 更新密码
	err = dao.UpdateUser(userForUpdate)
	if err != nil {
		global.Log.Error("用户密码更新失败", zap.Error(err))
		return nil, e.ErrorUpdateUser, e.IsLogicError
	}

	// 响应信息
	resp := &response.UserPasswordChangeResp{}

	return resp, e.Success, e.NotLogicError
}

func (*userService) UploadAvatar(c *gin.Context, file *multipart.FileHeader) (*response.UploadAvatarResp, e.Code, bool) {
	// 检查文件类型
	allowedTypes := []string{enum.AvatarTypeJPEG, enum.AvatarTypePNG}
	if !files.IsAllowedFileType(file, allowedTypes) {
		return nil, e.ErrorFileType, e.NotLogicError
	}

	// 从context中获取UserID
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	// 从context中获取Username
	username, ok := getUsername(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	// 保存文件
	ext := path.Ext(file.Filename)
	filename := strings.Join([]string{username, ext}, "")
	uploadMode := global.Config.Server.UploadMode
	url := getAvatarURL(uploadMode, filename)
	if uploadMode == enum.UploadModeLocal {
		dst := strings.Join([]string{".", global.Config.Image.AvatarDir, filename}, "/")
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			global.Log.Error("文件上传至本地错误", zap.Error(err))
			return nil, e.ErrorUploadFile, e.IsLogicError
		}
	} else {
		f, err := file.Open()
		if err != nil {
			global.Log.Error("文件内部错误", zap.Error(err))
			return nil, e.ErrorFileError, e.IsLogicError
		}
		err = oss.UploadFile(filename, f)
		if err != nil {
			global.Log.Error("文件上传至OSS错误", zap.Error(err))
			return nil, e.ErrorOSSUploadError, e.IsLogicError
		}
	}

	// 更新数据库
	userForUpdate := &model.User{
		Model:  gorm.Model{ID: userID},
		Avatar: url,
	}
	err := dao.UpdateUser(userForUpdate)
	if err != nil {
		global.Log.Error("用户头像上传失败", zap.Error(err))
		return nil, e.ErrorUpdateUser, e.IsLogicError
	}

	// 响应信息
	resp := &response.UploadAvatarResp{}

	return resp, e.Success, e.NotLogicError
}
