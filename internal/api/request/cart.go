package request

import "github.com/hespecial/gin-mall/pkg/validator"

type GetCartListReq struct{}

type AddCartItemReq struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func (*AddCartItemReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		// ...
	}
}

type UpdateCartItemQuantityReq struct {
	CartItemID uint `json:"id"`
	Quantity   int  `json:"quantity"`
}

func (*UpdateCartItemQuantityReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		// ...
	}
}

type DeleteCartItemReq struct {
	CartItemID uint `uri:"id"`
}

func (*DeleteCartItemReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		// ...
	}
}

type ClearCartReq struct{}

func (*ClearCartReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		// ...
	}
}
