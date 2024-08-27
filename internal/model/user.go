package model

import (
	"github.com/hespecial/gin-mall/pkg/encryption"
	"gorm.io/gorm"
)

const (
	DefaultUserStatus     = "Active"
	DefaultUserMoney      = "0"
	DefaultAvatarFileName = "default.jpeg"
)

type User struct {
	gorm.Model
	Username  string `gorm:"unique"`
	Email     string
	Password  string
	Nickname  string
	Status    string
	Avatar    string `gorm:"size:1000"`
	Money     string
	Relations []User    `gorm:"many2many:relation;"`
	Favorites []Product `gorm:"many2many:favorite;"`
}

func (u *User) GetUserID() uint {
	return u.ID
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) EncryptPassword() error {
	hash, err := encryption.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}

func (u *User) CheckPassword(password string) bool {
	return encryption.CheckPasswordHash(u.Password, password)
}

func (u *User) EncryptMoney() error {
	money, err := encryption.EncryptAES(u.Money)
	if err != nil {
		return err
	}
	u.Money = money
	return nil
}

func (u *User) DecryptMoney() error {
	money, err := encryption.DecryptAES(u.Money)
	if err != nil {
		return err
	}
	u.Money = money
	return nil
}
