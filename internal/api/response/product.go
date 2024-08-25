package response

type Product struct {
	ID       uint    `json:"id"`        // 商品ID
	Title    string  `json:"title"`     // 商品名称
	Price    float64 `json:"price"`     // 商品价格
	Stock    int     `json:"stock"`     // 库存数量
	ImageURL string  `json:"image_url"` // 商品主图的URL
}

type GetProductListResp struct {
	List  []*Product `json:"list"`
	Total int64      `json:"total"`
}

type ProductImage struct {
	ID       uint   `json:"id"` // 可选，视需求而定
	ImageURL string `json:"image_url"`
}

type GetProductDetailInfoResp struct {
	ID       uint            `json:"id"`
	Title    string          `json:"title"`
	Price    float64         `json:"price"`
	Stock    int             `json:"stock"`
	Category *Category       `json:"category"`
	Images   []*ProductImage `json:"images"`
}
