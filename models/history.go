package model

import (
	"github.com/jinzhu/gorm"
	"mwxy_create_articles/pkg/util"
)

type History struct {
	gorm.Model
	Name       string `json:"name"`
	Body       string `json:"body" gorm:"type:text"`
	ArticleIDS string `json:"article_ids"`
}

//展示分类模型
type ShowHistory struct {
	Name       string `json:"name"`
	Body       string `json:"body"`
	ArticleIDS string `json:"article_ids"`
}

//创建历史
func (history *History) Create() (uint, error) {
	if err := DB.Create(history).Error; err != nil {
		util.Log().Warning("无法成功创建分类, %s", err)
	}
	return history.ID, nil
}

//删除历史
func DeleteHistoryByID(id uint) error {
	result := DB.Where("id = ?", id).Delete(&History{})
	return result.Error
}

//查找历史
func GetHistoryByID(id uint) (History, error) {
	var history History
	result := DB.First(&history, id)
	return history, result.Error
}

//列出历史
func ListHistory(page, pageSize int, order string) ([]History, int) {
	var (
		histories []History
		total     int
	)
	dbChain := DB
	//计算总数
	dbChain.Model(&History{}).Count(&total)

	//查询记录
	dbChain.Limit(pageSize).Offset((page - 1) * pageSize).Order(order).Find(&histories)
	return histories, total
}
