package explorer

import (
	"github.com/gin-gonic/gin"
	model "mwxy_create_articles/models"
	"mwxy_create_articles/pkg/serializer"
)

//创建 历史记录
type CreateHistoryService struct {
	Name       string `json:"name" binding:"required,min=1,max=255"`
	ArticleIDS string `json:"article_ids" binding:"required"`
	Body       string `json:"body"`
}

type GetHistoryByIDService struct {
	ID uint `json:"id" uri:"id" binding:"required"`
}

// 列出历史列表
type HistoryListService struct {
	Page    uint   `form:"page" binding:"required,min=1"`
	OrderBy string `form:"order_by" binding:"required,eq=created_at"`
	Order   string `form:"order" binding:"required,eq=DESC|eq=ASC"`
}

//创建
func (service *CreateHistoryService) Create(c *gin.Context) serializer.Response {
	//查询获取到的ids
	//ids := strings.Split(service.ArticleIDS, ",")
	//body, err := model.GetArticleBodyByIDS(ids)
	//if err != nil || len(body) == 0 {
	//	return serializer.Err(serializer.CodeNotFound, "没有找到数据", err)
	//}
	history := model.History{
		Name:       service.Name,
		ArticleIDS: service.ArticleIDS,
		Body:       service.Body,
	}
	id, err := history.Create()
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "创建失败", err)
	}
	return serializer.Response{
		Data: id,
	}
}

//删除
func (service *GetHistoryByIDService) Delete(c *gin.Context) serializer.Response {
	if err := model.DeleteHistoryByID(service.ID); err != nil {
		return serializer.Err(serializer.CodeDBError, "删除失败", err)
	}
	return serializer.Response{}
}

//查找单个
func (service *GetHistoryByIDService) GetHistoryByID(c *gin.Context) serializer.Response {
	history, err := model.GetHistoryByID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeNotFound, "未找到", err)
	}
	return serializer.BuildHistoryResponse(history)
}

//历史列表
func (service *HistoryListService) List(c *gin.Context) serializer.Response {
	histories, total := model.ListHistory(int(service.Page), 10, service.OrderBy+" "+
		service.Order)
	return serializer.BuildHistoryListRes(histories, total)
}
