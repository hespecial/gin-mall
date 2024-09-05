package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/internal/api/request"
	"github.com/hespecial/gin-mall/internal/common"
	"github.com/hespecial/gin-mall/internal/common/e"
	"github.com/hespecial/gin-mall/internal/service"
	"github.com/hespecial/gin-mall/pkg/validator"
)

// GetOrderList godoc
//
//	@Summary		获取订单列表
//	@Description	获取用户的所有订单
//	@Tags			Order
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Produce		json
//	@Success		200	{object}	common.Response{data=response.GetOrderListResp}
//	@Router			/order [get]
func GetOrderList(c *gin.Context) {
	var req request.GetOrderListReq

	resp, code, isLogicError := service.OrderService.GetOrderList(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// GetOrderInfo godoc
//
//	@Summary		获取订单详情
//	@Description	通过订单id获取订单信息
//	@Tags			Order
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Produce		json
//	@Param			id	path		uint	true	"订单id"
//	@Success		200	{object}	common.Response{data=response.GetOrderInfoResp}
//	@Router			/order/{id} [get]
func GetOrderInfo(c *gin.Context) {
	var req request.GetOrderInfoReq
	if err := c.ShouldBindUri(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.OrderService.GetOrderInfo(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// CreateOrder godoc
//
//	@Summary		创建订单
//	@Description	通过用户地址和订单项来创建订单
//	@Tags			Order
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Accept			json
//	@Produce		json
//	@Param			order	body	string	true	"地址id（address_id）、购物项(items){商品id（product_id）、数量(quantity)}"
//	@Success		200		{object}	common.Response{data=response.AddAddressResp}
//	@Router			/order [post]
func CreateOrder(c *gin.Context) {
	var req request.CreateOrderReq
	if err := c.ShouldBind(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.OrderService.CreateOrder(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// DeleteOrder godoc
//
//	@Summary		删除订单
//	@Description	通过订单id删除指定的订单
//	@Tags			Order
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Produce		json
//	@Param			id	path		uint	true	"订单id"
//	@Success		200	{object}	common.Response{data=response.DeleteOrderResp}
//	@Router			/order/{id} [delete]
func DeleteOrder(c *gin.Context) {
	var req request.DeleteOrderReq
	if err := c.ShouldBindUri(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.OrderService.DeleteOrder(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}
