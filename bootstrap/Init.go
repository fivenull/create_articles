package bootstrap

import (
	model "mwxy_create_articles/models"
	"mwxy_create_articles/pkg/cache"
	"mwxy_create_articles/pkg/conf"

	"github.com/gin-gonic/gin"
)

func Init(path string) {
	InitApplication()
	conf.Init(path)
	// Debug 关闭时，切换为生产模式
	if !conf.SystemConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	//缓存初始化
	cache.Init()
	if conf.SystemConfig.Mode == "master" {
		model.Init()
		InitStatic()
	}
}
