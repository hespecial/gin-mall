package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/internal/api/request"
	"github.com/hespecial/gin-mall/internal/common"
	"github.com/hespecial/gin-mall/internal/common/e"
	"github.com/hespecial/gin-mall/internal/service"
	"github.com/hespecial/gin-mall/pkg/validator"
)

// GetAddressList godoc
//
//	@Summary		获取用户地址列表
//	@Description	获取用户的所有地址信息
//	@Tags			Address
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Produce		json
//	@Success		200	{object}	common.Response{data=response.GetAddressListResp}
//	@Router			/address [get]
func GetAddressList(c *gin.Context) {
	var req request.GetAddressListReq

	resp, code, isLogicError := service.AddressService.GetAddressList(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// GetAddressInfo godoc
//
//	@Summary		获取地址信息
//	@Description	通过地址id获取地址信息
//	@Tags			Address
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Produce		json
//	@Param			id	path		uint	true	"地址id"
//	@Success		200	{object}	common.Response{data=response.GetAddressInfoResp}
//	@Router			/address/{id} [get]
func GetAddressInfo(c *gin.Context) {
	var req request.GetAddressInfoReq
	if err := c.ShouldBindUri(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.AddressService.GetAddressInfo(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// AddAddress godoc
//
//	@Summary		添加用户地址
//	@Description	创建一个新的用户地址，地址信息包括姓名、电话、地址
//	@Tags			Address
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			name	formData	string	true	"姓名"
//	@Param			phone	formData	string	true	"电话"
//	@Param			address	formData	string	true	"地址"
//	@Success		200		{object}	common.Response{data=response.AddAddressResp}
//	@Router			/address [post]
func AddAddress(c *gin.Context) {
	var req request.AddAddressReq
	if err := c.ShouldBind(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.AddressService.AddAddress(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// UpdateAddress godoc
//
//	@Summary		更新用户地址
//	@Description	通过指定地址id更新用户地址，可更新信息包括姓名、电话、地址
//	@Tags			Address
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			id		path		uint	true	"地址id"
//	@Param			name	formData	string	true	"姓名"
//	@Param			phone	formData	string	true	"电话"
//	@Param			address	formData	string	true	"地址"
//	@Success		200		{object}	common.Response{data=response.UpdateAddressResp}
//	@Router			/address/{id} [put]
func UpdateAddress(c *gin.Context) {
	var req request.UpdateAddressReq
	if err := c.ShouldBind(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}
	req.AddressID = c.Param("id")

	resp, code, isLogicError := service.AddressService.UpdateAddress(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// DeleteAddress godoc
//
//	@Summary		删除用户地址
//	@Description	通过地址id删除指定的地址
//	@Tags			Address
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Produce		json
//	@Param			id	path		uint	true	"地址id"
//	@Success		200	{object}	common.Response{data=response.DeleteAddressResp}
//	@Router			/address/{id} [delete]
func DeleteAddress(c *gin.Context) {
	var req request.DeleteAddressReq
	if err := c.ShouldBindUri(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.AddressService.DeleteAddress(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}
