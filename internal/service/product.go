package service

import (
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/api/request"
	"github.com/hespecial/gin-mall/internal/api/response"
	"github.com/hespecial/gin-mall/internal/common/e"
	"github.com/hespecial/gin-mall/internal/repository/dao"
	"github.com/hespecial/gin-mall/internal/repository/es"
	"go.uber.org/zap"
)

type productService struct{}

var ProductService = new(productService)

func (*productService) GetProductList(_ *gin.Context, req *request.GetProductListReq) (*response.GetProductListResp, e.Code, bool) {
	products, count, err := dao.GetProductList(req.Page, req.Size)
	if err != nil {
		global.Log.Error("获取商品列表失败", zap.Error(err))
		return nil, e.ErrorGetProductList, e.IsLogicError
	}

	var list []*response.Product
	for _, product := range products {
		list = append(list, &response.Product{
			ID:       product.ID,
			Title:    product.Title,
			Price:    product.Price,
			Stock:    product.Stock,
			ImageURL: product.Images[0].URL,
		})
	}

	resp := &response.GetProductListResp{
		List:  list,
		Total: count,
		Page:  req.Page,
		Size:  req.Size,
	}

	return resp, e.Success, e.NotLogicError
}

func (*productService) GetProductDetailInfo(_ *gin.Context, req *request.GetProductDetailInfoReq) (*response.GetProductDetailInfoResp, e.Code, bool) {
	product, err := dao.GetProductById(req.ID)
	if err != nil {
		global.Log.Error("根据ID获取商品失败", zap.Error(err))
		return nil, e.ErrorGetProductByID, e.IsLogicError
	}

	var images []*response.ProductImage
	for _, img := range product.Images {
		images = append(images, &response.ProductImage{
			ID:       img.ID,
			ImageURL: img.URL,
		})
	}

	resp := &response.GetProductDetailInfoResp{
		ID:    product.ID,
		Title: product.Title,
		Price: product.Price,
		Stock: product.Stock,
		Category: &response.Category{
			ID:   product.Category.ID,
			Name: product.Category.CategoryName,
		},
		Images: images,
	}

	return resp, e.Success, e.NotLogicError
}

func (*productService) SearchProduct(_ *gin.Context, req *request.SearchProductReq) (*response.SearchProductResp, e.Code, bool) {
	// mysql搜索
	// products, count, err := dao.SearchProduct(req.Keyword, req.Page, req.Size)
	// es搜索
	products, count, err := es.SearchProduct(req.Keyword, req.Page, req.Size)
	if err != nil {
		global.Log.Error("搜索商品失败", zap.Error(err))
		return nil, e.ErrorSearchProduct, e.IsLogicError
	}

	var list []*response.Product
	for _, product := range products {
		list = append(list, &response.Product{
			ID:       product.ID,
			Title:    product.Title,
			Price:    product.Price,
			Stock:    product.Stock,
			ImageURL: product.Images[0].URL,
		})
	}

	resp := &response.SearchProductResp{
		List:  list,
		Total: count,
		Page:  req.Page,
		Size:  req.Size,
	}

	return resp, e.Success, e.NotLogicError
}
