package response

type CartItem struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	ImageURL  string  `json:"image_url"`
}

type GetCartListResp struct {
	Items []*CartItem `json:"items"`
}

type AddCartItemResp struct {
	CartItemID uint `json:"cart_item_id"`
}

type UpdateCartItemQuantityResp struct{}

type DeleteCartItemResp struct{}

type ClearCartResp struct{}
