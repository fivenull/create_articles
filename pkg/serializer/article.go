package serializer

import model "mwxy_create_articles/models"

type ArticleShowBody struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Body       string `json:"body"`
	OrderBy    uint   `json:"order_by"`
	CreateDate string `json:"create_date,omitempty"`
}

type Article struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	OrderBy    uint   `json:"order_by"`
	CreateDate string `json:"create_date,omitempty"`
}

type ArticleName struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

//构建文章列表
func buildArticleRes(articles []model.Article) []Article {
	res := make([]Article, 0, len(articles))
	for i := 0; i < len(articles); i++ {
		newArticle := Article{
			ID:         articles[i].ID,
			Name:       articles[i].Name,
			OrderBy:    articles[i].OrderBy,
			CreateDate: articles[i].CreatedAt.Format("2006-01-02 15:04:05"),
		}
		res = append(res, newArticle)
	}
	return res
}

//构建文章列表 带body
func buildArticleShowBodyRes(articles []model.Article) []ArticleShowBody {
	res := make([]ArticleShowBody, 0, len(articles))
	for i := 0; i < len(articles); i++ {
		newArticle := ArticleShowBody{
			ID:         articles[i].ID,
			Name:       articles[i].Name,
			Body:       articles[i].Body,
			OrderBy:    articles[i].OrderBy,
			CreateDate: articles[i].CreatedAt.Format("2006-01-02 15:04:05"),
		}
		res = append(res, newArticle)
	}
	return res
}

// 序列化文章
func BuildArticleShowBody(article model.Article) ArticleShowBody {
	return ArticleShowBody{
		ID:         article.ID,
		Name:       article.Name,
		Body:       article.Body,
		OrderBy:    article.OrderBy,
		CreateDate: article.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func BuildArticleShowBodyResponse(article model.Article) Response {
	return Response{
		Data: BuildArticleShowBody(article),
	}
}

// 序列化文章列表
func buildArticleShowBodyList(articles []model.Article) []ArticleShowBody {
	res := make([]ArticleShowBody, 0, len(articles))
	for i := 0; i < len(articles); i++ {
		item := ArticleShowBody{
			ID:         articles[i].ID,
			Name:       articles[i].Name,
			Body:       articles[i].Body,
			OrderBy:    articles[i].OrderBy,
			CreateDate: articles[i].CreatedAt.Format("2006-01-02 15:04:05"),
		}
		res = append(res, item)
	}
	return res
}

//返回序列化文章列表，通过分类查询
func BuildArticleListByCategoryID(category model.Category, articles []model.Article) Response {
	cate := BuildCategory(category)
	return Response{Data: map[string]interface{}{
		"category": cate,
		"items":    buildArticleShowBodyList(articles),
	}}
}

//返回序列化文章列表，通过分类查询
func BuildArticleListName(articles []model.ArticleName) Response {
	res := make([]ArticleName, 0, len(articles))
	for i := 0; i < len(articles); i++ {
		item := ArticleName{
			ID:   articles[i].ID,
			Name: articles[i].Name,
		}
		res = append(res, item)
	}
	return Response{Data: map[string]interface{}{
		"items": res,
	}}
}

func BuildArticleShowBodyListByIDS(articles []model.Article) Response {
	return Response{Data: map[string]interface{}{
		"items": buildArticleShowBodyList(articles),
	}}
}
