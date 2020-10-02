package controllers

import (
	"github.com/gin-gonic/gin"
	"mwxy_create_articles/service/explorer"
)

//创建历史
func CreateHistory(c *gin.Context) {
	var service explorer.CreateHistoryService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//删除历史
func DeleteHistory(c *gin.Context) {
	var service explorer.GetHistoryByIDService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.Delete(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//历史列表
func ListHistory(c *gin.Context) {
	var service explorer.HistoryListService
	if err := c.ShouldBindQuery(&service); err == nil {
		res := service.List(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, err)
	}
}

//获取单个历史
func GetHistory(c *gin.Context) {
	var service explorer.GetHistoryByIDService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.GetHistoryByID(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
