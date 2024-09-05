package response

import "time"

type OrderItem struct {
	OrderItemID uint    `json:"order_item_id"`
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
}

type Order struct {
	OrderID       uint         `json:"order_id"`
	OrderNumber   string       `json:"order_number"`
	PaymentStatus int          `json:"payment_status"`
	TotalAmount   float64      `json:"total_amount"`
	Address       string       `json:"address"`
	Items         []*OrderItem `json:"items"`
	CreatedAt     time.Time    `json:"created_at"`
}

type GetOrderListResp struct {
	Orders []*Order `json:"orders"`
}

type GetOrderInfoResp struct {
	Order *Order `json:"order"`
}

type CreateOrderResp struct {
	OrderNumber   string  `json:"order_number"`
	PaymentStatus int     `json:"payment_status"`
	TotalAmount   float64 `json:"total_amount"`
}

type DeleteOrderResp struct{}
