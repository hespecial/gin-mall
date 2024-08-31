package service

import (
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/api/request"
	"github.com/hespecial/gin-mall/internal/api/response"
	"github.com/hespecial/gin-mall/internal/common/e"
	"github.com/hespecial/gin-mall/internal/model"
	"github.com/hespecial/gin-mall/internal/repository/dao"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
)

type addressService struct{}

var AddressService = new(addressService)

func (*addressService) GetAddressList(c *gin.Context, _ *request.GetAddressListReq) (*response.GetAddressListResp, e.Code, bool) {
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	addresses, err := dao.GetAddressList(userID)
	if err != nil {
		global.Log.Error("获取用户地址列表失败", zap.Error(err))
		return nil, e.ErrorGetAddressList, e.IsLogicError
	}

	var list []*response.Address
	for _, address := range addresses {
		list = append(list, &response.Address{
			ID:      address.ID,
			Name:    address.Name,
			Phone:   address.Phone,
			Address: address.Address,
		})
	}

	resp := &response.GetAddressListResp{
		List: list,
	}

	return resp, e.Success, e.NotLogicError
}

func (*addressService) GetAddressInfo(_ *gin.Context, req *request.GetAddressInfoReq) (*response.GetAddressInfoResp, e.Code, bool) {
	addr, err := dao.GetAddressInfo(req.AddressID)
	if err != nil {
		global.Log.Error("获取地址信息失败", zap.Error(err))
		return nil, e.ErrorGetAddressInfo, e.IsLogicError
	}

	resp := &response.GetAddressInfoResp{
		Address: &response.Address{
			ID:      addr.ID,
			Name:    addr.Name,
			Phone:   addr.Phone,
			Address: addr.Address,
		},
	}

	return resp, e.Success, e.NotLogicError
}

func (*addressService) AddAddress(c *gin.Context, req *request.AddAddressReq) (*response.AddAddressResp, e.Code, bool) {
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	addr := &model.Address{
		UserID:  userID,
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
	}

	addressID, err := dao.AddAddress(addr)
	if err != nil {
		global.Log.Error("添加用户地址失败", zap.Error(err))
		return nil, e.ErrorAddAddress, e.IsLogicError
	}

	resp := &response.AddAddressResp{
		AddressID: addressID,
	}

	return resp, e.Success, e.NotLogicError
}

func (*addressService) UpdateAddress(_ *gin.Context, req *request.UpdateAddressReq) (*response.UpdateAddressResp, e.Code, bool) {
	addressID, _ := strconv.Atoi(req.AddressID)

	addr := &model.Address{
		Model:   gorm.Model{ID: uint(addressID)},
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
	}

	err := dao.UpdateAddress(addr)
	if err != nil {
		global.Log.Error("更新用户地址失败", zap.Error(err))
		return nil, e.ErrorUpdateAddress, e.IsLogicError
	}

	resp := &response.UpdateAddressResp{}

	return resp, e.Success, e.NotLogicError
}

func (*addressService) DeleteAddress(_ *gin.Context, req *request.DeleteAddressReq) (*response.DeleteAddressResp, e.Code, bool) {
	err := dao.DeleteAddress(req.AddressID)
	if err != nil {
		global.Log.Error("删除用户地址失败", zap.Error(err))
		return nil, e.ErrorDeleteAddress, e.IsLogicError
	}

	resp := &response.DeleteAddressResp{}

	return resp, e.Success, e.NotLogicError
}
