package dao

import (
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/model"
)

func GetCategoryList() (list []*model.Category, _ error) {
	return list, global.DB.Find(&list).Error
}
