package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/internal/api/request"
	"github.com/hespecial/gin-mall/internal/common"
	"github.com/hespecial/gin-mall/internal/common/e"
	"github.com/hespecial/gin-mall/internal/service"
)

// GetCategoryList godoc
//
//	@Summary		获取商品分类列表
//	@Description	获取所有的商品分类
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	common.Response{data=response.GetCategoryListResp}
//	@Router			/category [get]
func GetCategoryList(c *gin.Context) {
	var req *request.GetCategoryListReq

	resp, code, isLogicError := service.CategoryService.GetCategoryList(c, req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}
