package model

import (
	"gorm.io/gorm"
	"time"
)

const (
	PaymentStatusPending = 0
	PaymentStatusPaid    = 1
)

type Order struct {
	gorm.Model
	OrderNumber     string `gorm:"unique;not null"` // 订单编号
	UserID          uint
	TotalAmount     float64 `gorm:"not null"` // 订单总金额
	AddressID       uint
	Address         Address     `gorm:"foreignKey:AddressID"`
	Items           []OrderItem `gorm:"foreignKey:OrderID"` // 订单项
	PaymentMethod   string      // 支付方式
	PaymentStatus   int         `gorm:"not null"` // 支付状态 0 未支付 1 已支付
	PaymentDeadline time.Time   // 支付截至时间
}

type OrderItem struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID"` // 商品信息
	Quantity  int     `gorm:"not null"`             // 购买数量
}
