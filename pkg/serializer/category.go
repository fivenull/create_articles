package serializer

import (
	model "mwxy_create_articles/models"
)

// Category 序列化器
type Category struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Info       string `json:"info"`
	OrderBy    uint   `json:"order_by"`
	CreateDate string `json:"create_date,omitempty"`
}

// 序列化分类
func BuildCategory(category model.Category) Category {
	return Category{
		ID:         category.ID,
		Name:       category.Name,
		Info:       category.Info,
		OrderBy:    category.OrderBy,
		CreateDate: category.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func BuildCategoryResponse(category model.Category) Response {
	return Response{
		Data: BuildCategory(category),
	}
}

// 序列化分类列表
func BuildCategoryList(categories []model.Category) Response {
	res := make([]Category, 0, len(categories))
	for i := 0; i < len(categories); i++ {
		item := Category{
			ID:         categories[i].ID,
			Name:       categories[i].Name,
			Info:       categories[i].Info,
			OrderBy:    categories[i].OrderBy,
			CreateDate: categories[i].CreatedAt.Format("2006-01-02 15:04:05"),
		}
		res = append(res, item)
	}
	return Response{Data: map[string]interface{}{
		"items": res,
	}}
}
