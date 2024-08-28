package dao

import (
	"errors"
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/model"
	"gorm.io/gorm"
)

func createCart(userID uint) (cart *model.Cart, _ error) {
	cart = &model.Cart{
		UserID: userID,
	}
	return cart, global.DB.Create(&cart).Error
}

func cartNotExist(userID uint) bool {
	var cart model.Cart
	result := global.DB.Where("user_id = ?", userID).First(&cart)
	return errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func getCart(userID uint) (cart *model.Cart, _ error) {
	return cart, global.DB.Where("user_id = ?", userID).First(&cart).Error
}

func GertCartList(userID uint) (cart *model.Cart, err error) {
	if cartNotExist(userID) {
		cart, err = createCart(userID)
	} else {
		cart, err = getCart(userID)
	}
	if err != nil {
		return nil, err
	}

	err = global.DB.
		Model(&model.Cart{}).
		Where("user_id = ?", userID).
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.Product.Images").
		First(&cart).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}

	return
}

func AddCartItem(userID, productID uint, quantity int) (cartItemID uint, err error) {
	var cart *model.Cart
	if cartNotExist(userID) {
		cart, err = createCart(userID)
	} else {
		cart, err = getCart(userID)
	}
	if err != nil {
		return 0, err
	}

	cartItem := &model.CartItem{
		CartID:    cart.ID,
		ProductID: productID,
		Quantity:  quantity,
	}

	return cartItem.ID, global.DB.Create(&cartItem).Error
}

func UpdateCartItemQuantity(cartItemID uint, quantity int) (err error) {
	return global.DB.
		Model(&model.CartItem{}).
		Where("id = ?", cartItemID).
		Update("quantity", quantity).Error
}

func DeleteCartItem(cartItemID uint) (err error) {
	return global.DB.Delete(&model.CartItem{}, cartItemID).Error
}

func ClearCart(userID uint) (err error) {
	cart, err := getCart(userID)
	if err != nil {
		return err
	}
	return global.DB.
		Where("cart_id = ?", cart.ID).
		Delete(&model.CartItem{}).Error
}
