package dao

import (
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/model"
)

func ExistByUsername(username string) (*model.User, bool) {
	var user *model.User
	count := global.DB.Where("username = ?", username).First(&user).RowsAffected
	if count == 0 {
		return nil, false
	}
	return user, true
}

func GetUserByID(id uint) (*model.User, error) {
	var user *model.User
	err := global.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user *model.User) error {
	return global.DB.Create(user).Error
}

func UpdateUser(user *model.User) error {
	return global.DB.Updates(user).Error
}
