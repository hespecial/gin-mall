package service

import (
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/api/request"
	"github.com/hespecial/gin-mall/internal/api/response"
	"github.com/hespecial/gin-mall/internal/common/e"
	"github.com/hespecial/gin-mall/internal/repository/cache"
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

	resp := &response.GetCartListResp{}
	// 先查询缓存
	items, exist, err := cache.GetCartList(userID)
	if err != nil {
		global.Log.Error("redis获取购物车失败", zap.Error(err))
		return nil, e.ErrorGetCart, e.IsLogicError
	}
	// 存在则直接返回
	if exist {
		resp.Items = items
		return resp, e.Success, e.NotLogicError
	}

	// 否则查询数据库
	cart, err := dao.GetCartList(userID)
	if err != nil {
		global.Log.Error("mysql获取购物车失败", zap.Error(err))
		return nil, e.ErrorGetCart, e.IsLogicError
	}

	for _, item := range cart.Items {
		items = append(items, &response.CartItem{
			ID:        item.ID,
			ProductID: item.ProductID,
			Title:     item.Product.Title,
			Price:     item.Product.Price * float64(item.Quantity),
			Quantity:  item.Quantity,
			ImageURL:  item.Product.Images[0].URL,
		})
	}
	resp.Items = items

	// 写回缓存
	err = cache.SaveCartItems(userID, items)
	if err != nil {
		global.Log.Error("缓存购物项失败", zap.Error(err))
		return nil, e.ErrorCacheCartItems, e.IsLogicError
	}

	return resp, e.Success, e.NotLogicError
}

func (*cartService) AddCartItem(c *gin.Context, req *request.AddCartItemReq) (*response.AddCartItemResp, e.Code, bool) {
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	// 先更新数据库
	cartItemID, err := dao.AddCartItem(userID, req.ProductID, req.Quantity)
	if err != nil {
		global.Log.Error("添加购物项失败", zap.Error(err))
		return nil, e.ErrorAddCartItem, e.IsLogicError
	}

	// 再删除缓存
	err = cache.DeleteCartCache(userID)
	if err != nil {
		global.Log.Error("删除购物车缓存失败", zap.Error(err))
		return nil, e.ErrorDeleteCartCache, e.IsLogicError
	}

	resp := &response.AddCartItemResp{
		CartItemID: cartItemID,
	}

	return resp, e.Success, e.NotLogicError
}

func (*cartService) UpdateCartItemQuantity(c *gin.Context, req *request.UpdateCartItemQuantityReq) (*response.UpdateCartItemQuantityResp, e.Code, bool) {
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	// 先更新数据库
	err := dao.UpdateCartItemQuantity(req.CartItemID, req.Quantity)
	if err != nil {
		global.Log.Error("更新购物项数量失败", zap.Error(err))
		return nil, e.ErrorUpdateCartItemQuantity, e.IsLogicError
	}

	// 再删除缓存
	err = cache.DeleteCartCache(userID)
	if err != nil {
		global.Log.Error("删除购物车缓存失败", zap.Error(err))
		return nil, e.ErrorDeleteCartCache, e.IsLogicError
	}

	resp := &response.UpdateCartItemQuantityResp{}

	return resp, e.Success, e.NotLogicError
}

func (*cartService) DeleteCartItem(c *gin.Context, req *request.DeleteCartItemReq) (*response.DeleteCartItemResp, e.Code, bool) {
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	// 先更新数据库
	err := dao.DeleteCartItem(req.CartItemID)
	if err != nil {
		global.Log.Error("删除购物项失败", zap.Error(err))
		return nil, e.ErrorDeleteCartItem, e.IsLogicError
	}

	// 再删除缓存
	err = cache.DeleteCartCache(userID)
	if err != nil {
		global.Log.Error("删除购物车缓存失败", zap.Error(err))
		return nil, e.ErrorDeleteCartCache, e.IsLogicError
	}

	resp := &response.DeleteCartItemResp{}

	return resp, e.Success, e.NotLogicError
}

func (*cartService) ClearCart(c *gin.Context, _ *request.ClearCartReq) (*response.ClearCartResp, e.Code, bool) {
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	// 先更新数据库
	err := dao.ClearCart(userID)
	if err != nil {
		global.Log.Error("清空购物车失败", zap.Error(err))
		return nil, e.ErrorClearCart, e.IsLogicError
	}

	// 再删除缓存
	err = cache.DeleteCartCache(userID)
	if err != nil {
		global.Log.Error("删除购物车缓存失败", zap.Error(err))
		return nil, e.ErrorDeleteCartCache, e.IsLogicError
	}

	resp := &response.ClearCartResp{}

	return resp, e.Success, e.NotLogicError
}
