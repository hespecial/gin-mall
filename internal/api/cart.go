package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/internal/api/request"
	"github.com/hespecial/gin-mall/internal/common"
	"github.com/hespecial/gin-mall/internal/common/e"
	"github.com/hespecial/gin-mall/internal/service"
	"github.com/hespecial/gin-mall/pkg/validator"
)

// GetCartList godoc
//
//	@Summary		获取购物车列表
//	@Description	获取用户所有加入购物车的商品
//	@Tags			Cart
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Produce		json
//	@Success		200	{object}	common.Response{data=response.GetCartListResp}
//	@Router			/cart [get]
func GetCartList(c *gin.Context) {
	var req request.GetCartListReq

	resp, code, isLogicError := service.CartService.GetCartList(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// AddCartItem godoc
//
//	@Summary		添加购物项
//	@Description	将指定的商品和数量添加至购物车
//	@Tags			Cart
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Accept			json
//	@Produce		json
//	@Param			cart_item	body		string	true	"商品id（product_id）、数量(quantity)"
//	@Success		200			{object}	common.Response{data=response.AddCartItemResp}
//	@Router			/cart/item [post]
func AddCartItem(c *gin.Context) {
	var req request.AddCartItemReq
	if err := c.ShouldBind(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.CartService.AddCartItem(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// UpdateCartItemQuantity godoc
//
//	@Summary		更新购物项数量
//	@Description	更改购物车中指定的商品数量
//	@Tags			Cart
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Accept			json
//	@Produce		json
//	@Param			cart_item	body		string	true	"购物项id（id）、数量（quantity）"
//	@Success		200			{object}	common.Response{data=response.UpdateCartItemQuantityResp}
//	@Router			/cart/item [put]
func UpdateCartItemQuantity(c *gin.Context) {
	var req request.UpdateCartItemQuantityReq
	if err := c.ShouldBind(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.CartService.UpdateCartItemQuantity(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// DeleteCartItem godoc
//
//	@Summary		删除购物项
//	@Description	通过购物项id删除购物车中的商品
//	@Tags			Cart
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Produce		json
//	@Param			id	path		string	true	"购物项id"
//	@Success		200	{object}	common.Response{data=response.DeleteCartItemResp}
//	@Router			/cart/item/{id} [delete]
func DeleteCartItem(c *gin.Context) {
	var req request.DeleteCartItemReq
	if err := c.ShouldBindUri(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.CartService.DeleteCartItem(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// ClearCart godoc
//
//	@Summary		清空购物车
//	@Description	清空用户的购物车
//	@Tags			Cart
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Produce		json
//	@Success		200	{object}	common.Response{data=response.ClearCartResp}
//	@Router			/cart [delete]
func ClearCart(c *gin.Context) {
	var req request.ClearCartReq

	resp, code, isLogicError := service.CartService.ClearCart(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}
