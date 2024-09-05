package request

import "github.com/hespecial/gin-mall/pkg/validator"

type GetOrderListReq struct{}

type GetOrderInfoReq struct {
	OrderID uint `uri:"id"`
}

func (*GetOrderInfoReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		// ...
	}
}

type OrderItem struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

type CreateOrderReq struct {
	AddressID uint         `json:"address_id" binding:"required"`
	Items     []*OrderItem `json:"items" binding:"required"`
}

func (*CreateOrderReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		// ...
	}
}

type DeleteOrderReq struct {
	OrderID uint `uri:"id"`
}

func (*DeleteOrderReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		// ...
	}
}
