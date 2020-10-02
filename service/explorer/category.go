package explorer

import (
	"github.com/gin-gonic/gin"
	model "mwxy_create_articles/models"
	"mwxy_create_articles/pkg/serializer"
	"strconv"
)

// 创建文章分类
type CategoryCreateAndUpdateService struct {
	Name    string `json:"name" binding:"required,min=1,max=255"`
	Info    string `json:"info" binding:"required,min=1,max=255"`
	OrderBy uint   `json:"order_by" binding:"required"`
}

//展示文章分类
type CategoryListService struct {
	Page     int    `json:"page" binding:"min=1,required"`
	PageSize int    `json:"page_size" binding:"min=1,required"`
	OrderBy  uint   `json:"order_by"`
	Name     string `json:"name"`
	Info     string `json:"info"`
}

// NoParamService 无需参数的服务
type NoParamService struct {
}

//查询文章分类
type CategoryService struct {
	ID uint `uri:"id" json:"id" binding:"required"`
}

//func (service *CategoryListService) Lists() serializer.Response {
//	var res []model.Category
//	total := 0
//
//	tx := model.DB.Model(&model.Category{})
//	if service.OrderBy != "" {
//		tx = tx.Order(service.OrderBy)
//	}
//	// 计算总数用于分页
//	tx.Count(&total)
//
//	//查询记录
//	tx.Limit(service.PageSize).Offset((service.Page - 1) * service.PageSize).Find(&res)
//
//	return serializer.Response{Data: map[string]interface{}{
//		"total": total,
//		"items": res,
//	}}
//}
//获取列表
//获取总的
func (service *NoParamService) CategoryLists() serializer.Response {
	var res []model.Category
	model.DB.Model(&model.Category{}).Order("order_by desc").Find(&res)
	return serializer.BuildCategoryList(res)
}

//获取单个
func (service *CategoryService) Get() serializer.Response {
	category, err := model.GetCategoryByID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeNotFound, ""+
			""+
			"", err)
	}
	return serializer.BuildCategoryResponse(category)
}

//删除
func (service *CategoryService) Delete(c *gin.Context) serializer.Response {
	category, err := model.GetCategoryByID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeNotFound, "分类不存在", err)
	}

	if err := model.DeleteCategoryByID(category.ID); err != nil {
		return serializer.Err(serializer.CodeDBError, "删除失败", err)
	}

	return serializer.Response{}
}

//创建分类
func (service *CategoryCreateAndUpdateService) Create(c *gin.Context) serializer.Response {
	//创建
	category := model.Category{
		Name:    service.Name,
		Info:    service.Info,
		OrderBy: service.OrderBy,
	}
	id, err := category.Create()
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "分类创建失败", err)
	}
	return serializer.Response{
		Data: id,
	}
}

//更新分类
func (service *CategoryCreateAndUpdateService) Update(c *gin.Context) serializer.Response {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := model.GetCategoryByID(uint(id))
	if err != nil {
		return serializer.Err(serializer.CodeNotFound, "没有找到分类", err)
	}

	val := map[string]interface{}{
		"name":     service.Name,
		"info":     service.Info,
		"order_by": service.OrderBy,
	}
	if err := category.Update(val); err != nil {
		return serializer.DBErr("无法更新分类", err)
	}
	return serializer.Response{
		Msg: "更新成功",
	}
}
