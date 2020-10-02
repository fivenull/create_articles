package model

import "mwxy_create_articles/pkg/util"

func migration() {
	util.Log().Info("开始进行数据库初始化")
	DB.AutoMigrate(&Article{}, &Category{}, &History{})
	util.Log().Info("数据库初始化结束")
}
