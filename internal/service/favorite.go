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

type favoriteService struct{}

var FavoriteService = new(favoriteService)

func (*favoriteService) GetFavoriteList(c *gin.Context, _ *request.GetFavoriteListReq) (*response.GetFavoriteListResp, e.Code, bool) {
	// 从context中获取UserID
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	favorites, err := dao.GetFavoriteList(userID)
	if err != nil {
		global.Log.Error("获取收藏列表失败", zap.Error(err))
		return nil, e.ErrorGetFavoriteList, e.IsLogicError
	}

	var list []*response.Favorite
	for _, favorite := range favorites {
		list = append(list, &response.Favorite{
			ID:       favorite.ID,
			Title:    favorite.Title,
			Price:    favorite.Price,
			ImageURL: favorite.Images[0].URL,
		})
	}

	resp := &response.GetFavoriteListResp{
		List: list,
	}

	return resp, e.Success, e.NotLogicError
}

func (*favoriteService) AddFavorite(c *gin.Context, req *request.AddFavoriteReq) (*response.AddFavoriteResp, e.Code, bool) {
	// 从context中获取UserID
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	err := dao.AddFavorite(userID, req.ID)
	if err != nil {
		global.Log.Error("收藏商品失败", zap.Error(err))
		return nil, e.ErrorAddFavorite, e.IsLogicError
	}

	resp := &response.AddFavoriteResp{}

	return resp, e.Success, e.NotLogicError
}

func (*favoriteService) DeleteFavorite(c *gin.Context, req *request.DeleteFavoriteReq) (*response.DeleteFavoriteResp, e.Code, bool) {
	// 从context中获取UserID
	userID, ok := getUserID(c)
	if !ok {
		return nil, e.ErrorContextValue, e.IsLogicError
	}

	err := dao.DeleteFavorite(userID, req.ID)
	if err != nil {
		global.Log.Error("取消收藏失败", zap.Error(err))
		return nil, e.ErrorDeleteFavorite, e.IsLogicError
	}

	resp := &response.DeleteFavoriteResp{}

	return resp, e.Success, e.NotLogicError
}
