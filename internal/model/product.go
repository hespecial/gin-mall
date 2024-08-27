package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title      string         `gorm:"size:255;not null"` // 商品标题
	Price      float64        `gorm:"not null"`          // 商品价格
	Stock      int            `gorm:"not null"`          // 库存数量
	CategoryID uint           // 外键，指向商品分类
	Category   Category       `gorm:"foreignKey:CategoryID"` // 商品分类，使用GORM的外键关联
	Images     []ProductImage `gorm:"foreignKey:ProductID"`  // 商品图片，一对多关系
}

type ProductImage struct {
	ID        uint   `gorm:"primaryKey"`
	ProductID uint   // 外键，指向所属商品
	URL       string `gorm:"size:255;not null"` // 图片URL
}
