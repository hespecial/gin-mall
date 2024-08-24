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

type categoryService struct{}

var CategoryService = new(categoryService)

func (*categoryService) GetCategoryList(_ *gin.Context, _ *request.GetCategoryListReq) (*response.GetCategoryListResp, e.Code, bool) {
	category, err := dao.GetCategoryList()
	if err != nil {
		global.Log.Error("获取商品分类失败", zap.Error(err))
		return nil, e.ErrorGetCategoryList, e.IsLogicError
	}

	var list []*response.Category
	for _, v := range category {
		list = append(list, &response.Category{
			ID:   v.ID,
			Name: v.CategoryName,
		})
	}

	resp := &response.GetCategoryListResp{
		List:  list,
		Total: len(list),
	}

	return resp, e.Success, e.NotLogicError
}
