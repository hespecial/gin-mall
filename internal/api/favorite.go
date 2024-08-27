package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/internal/api/request"
	"github.com/hespecial/gin-mall/internal/common"
	"github.com/hespecial/gin-mall/internal/common/e"
	"github.com/hespecial/gin-mall/internal/service"
	"github.com/hespecial/gin-mall/pkg/validator"
)

// GetFavoriteList godoc
//
//	@Summary		获取收藏列表
//	@Description	登陆用户后可查看已收藏的商品
//	@Tags			Favorite
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	common.Response{data=response.GetFavoriteListResp}
//	@Router			/favorite [get]
func GetFavoriteList(c *gin.Context) {
	var req request.GetFavoriteListReq

	resp, code, isLogicError := service.FavoriteService.GetFavoriteList(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// AddFavorite godoc
//
//	@Summary		取消收藏
//	@Description	通过商品id收藏商品
//	@Tags			Favorite
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Accept			json
//	@Produce		json
//	@Param			id	body		string	true	"商品id"
//	@Success		200	{object}	common.Response{data=response.AddFavoriteResp}
//	@Router			/favorite [post]
func AddFavorite(c *gin.Context) {
	var req request.AddFavoriteReq
	if err := c.ShouldBind(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.FavoriteService.AddFavorite(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// DeleteFavorite godoc
//
//	@Summary		收藏商品
//	@Description	通过商品id取消收藏
//	@Tags			Favorite
//	@Security		AccessToken
//	@Security		RefreshToken
//	@Accept			json
//	@Produce		json
//	@Param			id	body		string	true	"商品id"
//	@Success		200	{object}	common.Response{data=response.DeleteFavoriteResp}
//	@Router			/favorite [delete]
func DeleteFavorite(c *gin.Context) {
	var req request.DeleteFavoriteReq
	if err := c.ShouldBind(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.FavoriteService.DeleteFavorite(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}
