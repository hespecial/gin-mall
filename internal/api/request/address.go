package request

import "github.com/hespecial/gin-mall/pkg/validator"

type GetAddressListReq struct{}

type GetAddressInfoReq struct {
	AddressID uint `uri:"id"`
}

func (*GetAddressInfoReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		// ...
	}
}

type AddAddressReq struct {
	Name    string `form:"name" json:"name" binding:"required"`
	Phone   string `form:"phone" json:"phone" binding:"required"`
	Address string `form:"address" json:"address" binding:"required"`
}

func (*AddAddressReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		// ...
	}
}

type UpdateAddressReq struct {
	AddressID string `uri:"id"`
	Name      string `form:"name" json:"name" binding:"required"`
	Phone     string `form:"phone" json:"phone" binding:"required"`
	Address   string `form:"address" json:"address" binding:"required"`
}

func (*UpdateAddressReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		// ...
	}
}

type DeleteAddressReq struct {
	AddressID uint `uri:"id"`
}

func (*DeleteAddressReq) ErrorMessages() validator.ValidateMessages {
	return validator.ValidateMessages{
		// ...
	}
}
