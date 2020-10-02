package model

import (
	"github.com/jinzhu/gorm"
	"mwxy_create_articles/pkg/util"
)

//分类模型
type Category struct {
	gorm.Model
	Name    string `json:"name"`
	Info    string `json:"info"`
	OrderBy uint   `json:"order_by"`
}

//分类模型-展示
type ShowCategory struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Info    string `json:"info"`
	OrderBy uint   `json:"order_by"`
}

//GetCategoryByID 用id 获取模型
func GetCategoryByID(id uint) (Category, error) {
	var category Category
	result := DB.Where("id = ?", id).First(&category)
	return category, result.Error
}

//Create 创建分类
func (category *Category) Create() (uint, error) {
	if err := DB.Create(category).Error; err != nil {
		util.Log().Warning("无法成功创建分类, %s", err)
	}
	return category.ID, nil
}

//DeleteCategoryByID 根据给定的ID来删除
func DeleteCategoryByID(id uint) error {
	result := DB.Where("id = ?", id).Delete(&Category{})
	return result.Error
}

//Update 更新数据
func (category *Category) Update(val map[string]interface{}) error {
	return DB.Model(category).Updates(val).Error
}

//检验是否存在
func CheckCategoryIsExists(id uint) (bool, error) {
	_, err := GetCategoryByID(id)
	if err == nil {
		return true, nil
	} else {
		return false, err
	}
}
