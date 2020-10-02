package model

import (
	"github.com/jinzhu/gorm"
	"mwxy_create_articles/pkg/util"
	"strconv"
	"strings"
)

type Article struct {
	gorm.Model
	CID      uint      `json:"c_id" gorm:"default:0"`
	Category *Category `gorm:"PRELOAD:false,association_autoupdate:false"`
	Name     string    `json:"name"`
	Body     string    `json:"body" gorm:"type:text"`
	OrderBy  uint      `json:"order_by"`
}

type ArticleName struct {
	gorm.Model
	CID      uint      `json:"c_id" gorm:"default:0"`
	Category *Category `gorm:"PRELOAD:false,association_autoupdate:false"`
	Name     string    `json:"name"`
	OrderBy  uint      `json:"order_by"`
}

//GetArticleByID 通过ID 获取文章
func GetArticleByID(ID interface{}) (Article, error) {
	var article Article
	result := DB.First(&article, ID)
	return article, result.Error
}

//GetArticlesByCID 通过CID 获取文章
func GetArticlesByCID(cid uint) ([]Article, error) {
	var articles []Article
	result := DB.Where("c_id = ?", cid).Order("order_by desc").Find(&articles)
	return articles, result.Error
}

//DeleteArticleByID 根据给定的ID来删除
func DeleteArticleByID(id uint) error {
	result := DB.Where("id = ?", id).Delete(&Article{})
	return result.Error
}

//Update 更新内容数据
func (article *Article) Update(val map[string]interface{}) error {
	return DB.Model(article).Updates(val).Error
}

// Create 创建内容数据
func (article *Article) Create() (uint, error) {
	if err := DB.Create(article).Error; err != nil {
		util.Log().Warning("无法创建内容， %s", err)
		return 0, err
	}
	return article.ID, nil
}

//通过 ids 获取内容，返回
func GetArticleBodyByIDS(ids []string) (string, error) {
	var articles []Article
	err := DB.Where("id in (?)", ids).Find(&articles).Error
	if err != nil {
		return "", err
	}
	//需要考虑重复id的情况
	oldArticle := make(map[uint]string)
	for _, article := range articles {
		oldArticle[article.ID] = article.Body
	}
	body := make([]string, 0, len(articles))
	//for index, article := range articles {
	//	fmt.Println(reflect.TypeOf(index))
	//	body = append(body, article.Body)
	//}
	for _, value := range ids {
		value, _ := strconv.Atoi(value)
		body = append(body, oldArticle[uint(value)])
	}
	bodyString := strings.Join(body, "\n")
	return bodyString, nil
}

// 获取所有的id和name
func GetArticleListName() ([]ArticleName, error) {
	var articles []ArticleName
	res := DB.Raw("select * from articles as a inner join categories as c on a.c_id=c.id and  a.deleted_at is null and  c.deleted_at is null").Order("a.order_by desc").Scan(&articles)
	return articles, res.Error
}
