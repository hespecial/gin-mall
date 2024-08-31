package dao

import (
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/model"
)

func GetAddressList(userID uint) (list []*model.Address, _ error) {
	return list, global.DB.
		Find(&list, "user_id = ?", userID).Error
}

func GetAddressInfo(addressID uint) (addr *model.Address, _ error) {
	return addr, global.DB.
		First(&addr, addressID).Error
}

func AddAddress(addr *model.Address) (addressID uint, _ error) {
	return addr.ID, global.DB.Create(&addr).Error
}

func UpdateAddress(addr *model.Address) error {
	return global.DB.Updates(&addr).Error
}

func DeleteAddress(addressID uint) error {
	return global.DB.Delete(&model.Address{}, addressID).Error
}
