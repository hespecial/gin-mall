package service

import (
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/api/request"
	"github.com/hespecial/gin-mall/internal/api/response"
	"github.com/hespecial/gin-mall/internal/common/e"
	"github.com/hespecial/gin-mall/internal/repository/dao"
	"go.uber.org/zap"
)

type cartService struct{}

var CartService = new(cartService)

func (*cartService) GetCartList(c *gin.Context, _ *request.GetCartListReq) (*response.GetCartListResp, e.Code, bool) {
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	cart, err := dao.GertCartList(userID)
	if err != nil {
		global.Log.Error("获取购物车失败", zap.Error(err))
		return nil, e.ErrorGetCart, e.IsLogicError
	}

	var items []*response.CartItem
	for _, item := range cart.Items {
		items = append(items, &response.CartItem{
			ProductID: item.ProductID,
			Title:     item.Product.Title,
			Price:     item.Product.Price * float64(item.Quantity),
			Quantity:  item.Quantity,
			ImageURL:  item.Product.Images[0].URL,
		})
	}

	resp := &response.GetCartListResp{
		Items: items,
	}

	return resp, e.Success, e.NotLogicError
}

func (*cartService) AddCartItem(c *gin.Context, req *request.AddCartItemReq) (*response.AddCartItemResp, e.Code, bool) {
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	cartItemID, err := dao.AddCartItem(userID, req.ProductID, req.Quantity)
	if err != nil {
		global.Log.Error("添加购物项失败", zap.Error(err))
		return nil, e.ErrorAddCartItem, e.IsLogicError
	}

	resp := &response.AddCartItemResp{
		CartItemID: cartItemID,
	}

	return resp, e.Success, e.NotLogicError
}

func (*cartService) UpdateCartItemQuantity(_ *gin.Context, req *request.UpdateCartItemQuantityReq) (*response.UpdateCartItemQuantityResp, e.Code, bool) {
	err := dao.UpdateCartItemQuantity(req.CartItemID, req.Quantity)
	if err != nil {
		global.Log.Error("更新购物项数量失败", zap.Error(err))
		return nil, e.ErrorUpdateCartItemQuantity, e.IsLogicError
	}

	resp := &response.UpdateCartItemQuantityResp{}

	return resp, e.Success, e.NotLogicError
}

func (*cartService) DeleteCartItem(_ *gin.Context, req *request.DeleteCartItemReq) (*response.DeleteCartItemResp, e.Code, bool) {
	err := dao.DeleteCartItem(req.CartItemID)
	if err != nil {
		global.Log.Error("删除购物项失败", zap.Error(err))
		return nil, e.ErrorDeleteCartItem, e.IsLogicError
	}

	resp := &response.DeleteCartItemResp{}

	return resp, e.Success, e.NotLogicError
}

func (*cartService) ClearCart(c *gin.Context, _ *request.ClearCartReq) (*response.ClearCartResp, e.Code, bool) {
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	err := dao.ClearCart(userID)
	if err != nil {
		global.Log.Error("清空购物车失败", zap.Error(err))
		return nil, e.ErrorClearCart, e.IsLogicError
	}

	resp := &response.ClearCartResp{}

	return resp, e.Success, e.NotLogicError
}
