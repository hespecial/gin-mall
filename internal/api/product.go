package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/internal/api/request"
	"github.com/hespecial/gin-mall/internal/common"
	"github.com/hespecial/gin-mall/internal/common/e"
	"github.com/hespecial/gin-mall/internal/service"
	"github.com/hespecial/gin-mall/pkg/validator"
)

// GetProductList godoc
//
//	@Summary		获取商品列表
//	@Description	通过页号和大小获取指定的商品列表
//	@Tags			Product
//	@Produce		json
//	@Param			page	query		int	true	"页号"
//	@Param			size	query		int	true	"大小"
//	@Success		200		{object}	common.Response{data=response.GetProductListResp}
//	@Router			/products [get]
func GetProductList(c *gin.Context) {
	var req request.GetProductListReq
	if err := c.ShouldBind(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.ProductService.GetProductList(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// GetProductDetailInfo godoc
//
//	@Summary		获取商品详情
//	@Description	通过商品ID获取商品详情信息
//	@Tags			Product
//	@Produce		json
//	@Param			id	path		uint	true	"商品id"
//	@Success		200	{object}	common.Response{data=response.GetProductListResp}
//	@Router			/product/{id} [get]
func GetProductDetailInfo(c *gin.Context) {
	var req request.GetProductDetailInfoReq
	if err := c.ShouldBindUri(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.ProductService.GetProductDetailInfo(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}

// SearchProduct godoc
//
//	@Summary		搜索商品
//	@Description	通过关键词并指定页号和大小来获取查询的商品
//	@Tags			Product
//	@Produce		json
//	@Param			keyword	query		string	true	"关键词"
//	@Param			page	query		int		true	"页号"
//	@Param			size	query		int		true	"大小"
//	@Success		200		{object}	common.Response{data=response.SearchProductResp}
//	@Router			/product/search [get]
func SearchProduct(c *gin.Context) {
	var req request.SearchProductReq
	if err := c.ShouldBind(&req); err != nil {
		common.FailWithMsg(c, e.InvalidParams, validator.GetErrorMsg(req, err))
		return
	}

	resp, code, isLogicError := service.ProductService.SearchProduct(c, &req)
	if code != e.Success {
		common.Fail(c, code, isLogicError)
		return
	}

	common.Success(c, resp)
}
