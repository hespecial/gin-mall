package dao

import (
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/model"
	"gorm.io/gorm"
)

func GetProductList(page int, size int) (products []*model.Product, count int64, _ error) {
	return products, count,
		global.DB.Transaction(func(tx *gorm.DB) error {
			err := tx.Model(&model.Product{}).Count(&count).Error
			if err != nil {
				return err
			}
			return global.DB.Model(&model.Product{}).
				Preload("Category").
				Preload("Images").
				Offset(size * (page - 1)).
				Limit(size).
				Find(&products).Error
		})
}

func GetProductById(id uint) (product *model.Product, _ error) {
	return product, global.DB.
		Preload("Category").
		Preload("Images").
		First(&product, id).Error
}

func GetProductsByIDs(ids []uint) (products []*model.Product, err error) {
	return products, global.DB.Where("id IN (?)", ids).Find(&products).Error
}

func SearchProduct(keyword string, page int, size int) (products []*model.Product, count int64, _ error) {
	return products, count, global.DB.Transaction(
		func(tx *gorm.DB) error {
			err := tx.Model(&model.Product{}).
				Where("title LIKE ?", "%"+keyword+"%").
				Count(&count).Error
			if err != nil {
				return err
			}
			return tx.Preload("Category").
				Preload("Images").
				Where("title LIKE ?", "%"+keyword+"%").
				Offset(size * (page - 1)).
				Limit(size).
				Find(&products).Error
		})
}
