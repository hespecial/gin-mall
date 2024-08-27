package dao

import (
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/model"
)

func GetFavoriteList(userID uint) (favorites []*model.Product, err error) {
	user, err := GetUserByID(userID)
	if err != nil {
		return
	}
	return favorites, global.DB.
		Model(&user).
		Preload("Images").
		Association("Favorites").
		Find(&favorites)
}

func AddFavorite(userID, productID uint) (err error) {
	user, err := GetUserByID(userID)
	if err != nil {
		return
	}
	product, err := GetProductById(productID)
	if err != nil {
		return
	}
	return global.DB.Model(&user).
		Association("Favorites").
		Append([]*model.Product{product})
}

func DeleteFavorite(userID, productID uint) (err error) {
	user, err := GetUserByID(userID)
	if err != nil {
		return
	}
	product, err := GetProductById(productID)
	if err != nil {
		return
	}
	return global.DB.Model(&user).
		Association("Favorites").
		Delete(&product)
}
