package routers

import (
	"mwxy_create_articles/bootstrap"
	"mwxy_create_articles/pkg/conf"
	"mwxy_create_articles/pkg/util"
	"mwxy_create_articles/routers/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	util.Log().Info("当前运行模式：Master")
	return InitMasterRouter()

}

// InitMasterRouter 初始化主机模式路由
func InitMasterRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/api/"})))
	// r.Use(middleware.InjectSiteInfo())
	r.Use(static.Serve("/", bootstrap.StaticFS))
	// r.GET("manifest.json", controllers.Manifest)

	InitCORS(r)
	v3 := r.Group("/api/v1/")
	{
		category := v3.Group("category")
		{
			//获取分类列表
			category.GET("lists", controllers.GetCategoryLists)
			//获取单个分类
			category.GET("get/:id", controllers.GetCategory)
			//新增分类
			category.POST("create", controllers.CreateCategory)
			//更新分类
			category.PUT(":id", controllers.UpdateCategory)
			//删除分类
			category.DELETE(":id", controllers.DeleteCategory)
		}
		article := v3.Group("article")
		{
			//新增文章
			article.POST("create", controllers.CreateArticle)
			//获取单篇文章
			article.GET("get/:id", controllers.GetArticle)
			//通过分类获取文章
			article.GET("category/:id", controllers.GetArticleListsByCategoryID)
			//更新文章
			article.PUT(":id", controllers.UpdateArticleByID)
			//删除文章
			article.DELETE(":id", controllers.DeleteArticleByID)
			//文章名称获取
			article.GET("name", controllers.GetArticleListName)
			//通过多个文章id 获取文章
			article.POST("ids", controllers.GetArticleBodyByIDS)
		}
		history := v3.Group("history")
		{
			//新增记录
			history.POST("create", controllers.CreateHistory)
			//删除记录
			history.DELETE(":id", controllers.DeleteHistory)
			//列表
			history.GET("", controllers.ListHistory)
			//获取单个记录
			history.GET("get/:id", controllers.GetHistory)
		}
	}

	return r
}

// InitCORS 初始化跨域配置
func InitCORS(router *gin.Engine) {
	if conf.CORSConfig.AllowOrigins[0] != "UNSET" {
		router.Use(cors.New(cors.Config{
			AllowOrigins:     conf.CORSConfig.AllowOrigins,
			AllowMethods:     conf.CORSConfig.AllowMethods,
			AllowHeaders:     conf.CORSConfig.AllowHeaders,
			AllowCredentials: conf.CORSConfig.AllowCredentials,
			ExposeHeaders:    conf.CORSConfig.ExposeHeaders,
		}))
		return
	}
}
