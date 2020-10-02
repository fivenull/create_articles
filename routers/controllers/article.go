package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mwxy_create_articles/service/explorer"
)

//创建文章
func CreateArticle(c *gin.Context) {
	var service explorer.ArticleCreateAndUpdateService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//获取单篇文章
func GetArticle(c *gin.Context) {
	var service explorer.GetArticleBIDService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.GetArticleByID(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//通过分类id获取文章列表
func GetArticleListsByCategoryID(c *gin.Context) {
	var service explorer.GetArticleByCategoryIDService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.GetArticleByCategoryID(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//删除文章
func DeleteArticleByID(c *gin.Context) {
	var service explorer.GetArticleBIDService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.DeleteArticleByID(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//更新文章
func UpdateArticleByID(c *gin.Context) {
	var service explorer.ArticleCreateAndUpdateService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.UpdateArticleByID(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//获取文章名称
func GetArticleListName(c *gin.Context) {
	var service explorer.NoParamService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.GetArticleNameList(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//通过多个id 获取内容
func GetArticleBodyByIDS(c *gin.Context) {
	var service explorer.ArticleIDSService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.GetArticleBodyByIDS(c)
		c.JSON(200, res)
	} else {
		fmt.Println(service, err)
		c.JSON(200, ErrorResponse(err))
	}
}
