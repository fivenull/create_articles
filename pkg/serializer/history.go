package serializer

import model "mwxy_create_articles/models"

type History struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Body       string `json:"body"`
	CreateDate string `json:"create_date"`
}

//构件历史列表
func BuildHistoryListRes(histories []model.History, total int) Response {
	res := make([]History, 0, total)
	for i := 0; i < len(histories); i++ {
		newHistory := History{
			ID:         histories[i].ID,
			Name:       histories[i].Name,
			Body:       histories[i].Body,
			CreateDate: histories[i].CreatedAt.Format("2006-01-02 15:04:05"),
		}
		res = append(res, newHistory)
	}

	return Response{Data: map[string]interface{}{
		"total": total,
		"items": res,
	}}
}

//返回单个历史
func BuildHistoryResponse(history model.History) Response {
	return Response{
		Data: History{
			ID:   history.ID,
			Name: history.Name,
			Body: history.Body,
		},
	}
}
