package dao

import (
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/model"
)

func GetOrderList(userID uint) (orders []*model.Order, err error) {
	return orders, global.DB.
		Model(&model.Order{}).
		Where("user_id = ?", userID).
		Preload("Address").
		Preload("Items").
		Preload("Items.Product").
		Find(&orders).Error
}

func GetOrderInfo(userID uint, orderID uint) (order *model.Order, err error) {
	return order, global.DB.
		Model(&model.Order{}).
		Where("user_id = ? AND id = ?", userID, orderID).
		Preload("Address").
		Preload("Items").
		Preload("Items.Product").
		First(&order).Error
}

func CreateOrder(order *model.Order) (err error) {
	return global.DB.
		Create(order).Error
}

func DeleteOrder(orderID uint) error {
	return global.DB.Delete(&model.Order{}, orderID).Error
}
