package dao

import (
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/model"
	"gorm.io/gorm"
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

func FollowUser(userID, followID uint) error {
	var user, follow *model.User
	return global.DB.Transaction(func(tx *gorm.DB) error {
		tx.Model(&model.User{}).Where("id = ?", userID).First(&user)
		tx.Model(&model.User{}).Where("id = ?", followID).First(&follow)
		return tx.Model(&user).Association("Relations").Append([]*model.User{follow})
	})
}

func UnfollowUser(userID, followID uint) error {
	var user, follow *model.User
	return global.DB.Transaction(func(tx *gorm.DB) error {
		tx.Model(&model.User{}).Where("id = ?", userID).First(&user)
		tx.Model(&model.User{}).Where("id = ?", followID).First(&follow)
		return tx.Model(&user).Association("Relations").Delete(&follow)
	})
}

func GetFollowingUsers(userID uint) ([]*model.User, error) {
	var user model.User
	var following []*model.User
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 查找用户
		if err := tx.First(&user, userID).Error; err != nil {
			return err
		}
		// 查找该关注对象
		if err := tx.Model(&user).Association("Relations").Find(&following); err != nil {
			return err
		}
		return nil
	})
	return following, err
}

func GetFollowerUsers(userID uint) ([]*model.User, error) {
	var followers []*model.User
	return followers, global.DB.Joins("JOIN relation ON relation.user_id = user.id").
		Where("relation.relation_id = ?", userID).
		Find(&followers).Error
}
