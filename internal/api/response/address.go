package response

type Address struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type GetAddressListResp struct {
	List []*Address `json:"list"`
}

type GetAddressInfoResp struct {
	Address *Address `json:"address"`
}

type AddAddressResp struct {
	AddressID uint `json:"address_id"`
}

type UpdateAddressResp struct{}

type DeleteAddressResp struct{}
