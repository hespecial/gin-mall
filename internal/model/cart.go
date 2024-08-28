package model

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID    uint
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int
}

type Cart struct {
	gorm.Model
	UserID uint
	Items  []CartItem `gorm:"foreignKey:CartID"`
}
