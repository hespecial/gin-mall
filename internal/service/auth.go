package service

import (
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/api/request"
	"github.com/hespecial/gin-mall/internal/api/response"
	"github.com/hespecial/gin-mall/internal/common/e"
	"github.com/hespecial/gin-mall/internal/model"
	"github.com/hespecial/gin-mall/internal/repository/dao"
	"github.com/hespecial/gin-mall/pkg/jwt"
	"github.com/hespecial/gin-mall/pkg/random"
	"go.uber.org/zap"
)

type authService struct{}

var AuthService = new(authService)

func (*authService) Register(_ *gin.Context, req *request.AuthRegisterReq) (*response.AuthRegisterResp, e.Code, bool) {
	// 检查用户名是否重复
	if _, exist := dao.ExistByUsername(req.Username); exist {
		global.Log.Info("用户名已存在")
		return nil, e.ErrorUserExists, e.NotLogicError
	}

	// user实例
	user := &model.User{
		Username: req.Username,
		Password: req.Password,
		Nickname: random.GenerateNickname(),
		Avatar:   getAvatarURL(global.Config.Server.UploadMode, model.DefaultAvatarFileName),
		Status:   model.DefaultUserStatus,
		Money:    model.DefaultUserMoney,
	}

	// 加密
	if err := user.EncryptPassword(); err != nil {
		global.Log.Error("密码加密错误", zap.Error(err))
		return nil, e.ErrorEncryptPassword, e.IsLogicError
	}
	if err := user.EncryptMoney(); err != nil {
		global.Log.Error("金额加密错误", zap.Error(err))
		return nil, e.ErrorEncryptMoney, e.IsLogicError
	}

	// 创建用户
	if err := dao.CreateUser(user); err != nil {
		global.Log.Error("创建用户失败", zap.Error(err))
		return nil, e.ErrorCreateUser, e.IsLogicError
	}

	// 响应信息
	resp := &response.AuthRegisterResp{}

	return resp, e.Success, e.NotLogicError
}

func (*authService) Login(_ *gin.Context, req *request.AuthLoginReq) (*response.AuthLoginResp, e.Code, bool) {
	// 检查用户名是否存在
	user, exist := dao.ExistByUsername(req.Username)
	if !exist {
		global.Log.Info("用户名不存在")
		return nil, e.ErrorAccountInvalid, e.NotLogicError
	}

	// 校验密码
	consist := user.CheckPassword(req.Password)
	if !consist {
		global.Log.Info("密码错误")
		return nil, e.ErrorAccountInvalid, e.NotLogicError
	}

	// 解密金额
	// err := user.DecryptMoney()
	// if err != nil {
	// 	global.Log.Error("金额解密错误", zap.Error(err))
	// 	return nil, e.ErrorDecryptMoney.Error()
	// }

	// 生成token
	accessToken, refreshToken, err := jwt.GenerateToken(user)
	if err != nil {
		global.Log.Error("生成token错误", zap.Error(err))
		return nil, e.ErrorGenerateToken, e.IsLogicError
	}

	// 响应数据
	resp := &response.AuthLoginResp{
		Nickname:     user.Nickname,
		Avatar:       user.Avatar,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return resp, e.Success, e.NotLogicError
}
