package explorer

import (
	"github.com/gin-gonic/gin"
	model "mwxy_create_articles/models"
	"mwxy_create_articles/pkg/serializer"
	"strconv"
	"strings"
)

// 创建文章 cid name body order
type ArticleCreateAndUpdateService struct {
	//Category model.Category `json:"Category" binding:"required"`
	CID     uint   `json:"c_id" binding:"required,min=1"`
	Name    string `json:"name" binding:"required,min=1,max=255"`
	Body    string `json:"body" binding:"required,min=1,max=255"`
	OrderBy uint   `json:"order_by" binding:"required"`
}

//查询文章
type ArticleService struct {
	ID uint `uri:"id" json:"id" binding:"required"`
}

//通过文章分类查询ID
type GetArticleByCategoryIDService struct {
	ID uint `uri:"id" json:"id" binding:"required"`
}

//通过文章ID查询
type GetArticleBIDService struct {
	ID uint `uri:"id" json:"id" binding:"required"`
}

type ArticleIDSService struct {
	ArticleIDS string `json:"article_ids" binding:"required"`
}

//创建文章
func (service *ArticleCreateAndUpdateService) Create(c *gin.Context) serializer.Response {
	_, err := model.GetCategoryByID(service.CID)
	if err != nil {
		return serializer.Err(serializer.CodeNotFound, "分类不存在", err)
	}
	article := model.Article{
		CID:     service.CID,
		Name:    service.Name,
		Body:    service.Body,
		OrderBy: service.OrderBy,
	}
	id, err := article.Create()
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "分类创建失败", err)
	}
	return serializer.Response{
		Data: id,
	}
}

//通过id获取单篇文章
func (service *GetArticleBIDService) GetArticleByID(c *gin.Context) serializer.Response {
	article, err := model.GetArticleByID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeNotFound, "文章不存在", err)
	}
	return serializer.BuildArticleShowBodyResponse(article)
}

//通过分类id 查询所有文章
func (service *GetArticleByCategoryIDService) GetArticleByCategoryID(c *gin.Context) serializer.Response {
	category, err := model.GetCategoryByID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeNotFound, "文章分类不存在", err)
	}
	articles, err := model.GetArticlesByCID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "", err)
	}
	return serializer.BuildArticleListByCategoryID(category, articles)
}

//通过id删除文章
func (service *GetArticleBIDService) DeleteArticleByID(c *gin.Context) serializer.Response {
	if err := model.DeleteArticleByID(service.ID); err != nil {
		return serializer.Err(serializer.CodeDBError, "删除失败", err)
	}
	return serializer.Response{}
}

//通过id更新文章
func (service *ArticleCreateAndUpdateService) UpdateArticleByID(c *gin.Context) serializer.Response {
	id, _ := strconv.Atoi(c.Param("id"))

	article, err := model.GetArticleByID(uint(id))
	if err != nil {
		return serializer.Err(serializer.CodeNotFound, "没有找到文章", err)
	}
	val := map[string]interface{}{
		"name":     service.Name,
		"c_id":     service.CID,
		"Body":     service.Body,
		"order_by": service.OrderBy,
	}
	if err := article.Update(val); err != nil {
		return serializer.DBErr("无法更新文章", err)
	}
	return serializer.Response{
		Msg: "更新成功",
	}
}

//文章名称列表
func (service *NoParamService) GetArticleNameList(c *gin.Context) serializer.Response {
	articles, err := model.GetArticleListName()
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "", err)
	}
	return serializer.BuildArticleListName(articles)
}

//通过id获取文章内容
func (service *ArticleIDSService) GetArticleBodyByIDS(c *gin.Context) serializer.Response {
	ids := strings.Split(service.ArticleIDS, ",")
	articles, err := model.GetArticleBodyByIDS(ids)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "", err)
	}
	return serializer.Response{
		Data: articles,
	}
}
