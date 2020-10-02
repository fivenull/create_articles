package controllers

import (
	"github.com/gin-gonic/gin"
	"mwxy_create_articles/service/explorer"
)

//创建分类
func CreateCategory(c *gin.Context) {
	var service explorer.CategoryCreateAndUpdateService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//删除分类
func DeleteCategory(c *gin.Context) {
	var service explorer.CategoryService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Delete(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//更新分类
func UpdateCategory(c *gin.Context) {
	var service explorer.CategoryCreateAndUpdateService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Update(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//获取单个分类
func GetCategory(c *gin.Context) {
	var service explorer.CategoryService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Get()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//获取全部分类
func GetCategoryLists(c *gin.Context) {
	var service explorer.NoParamService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.CategoryLists()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
