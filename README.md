# create_articles
gin+vue 入门案例，快速生成文章程序

gin 项目参照[cloudreve](https://github.com/cloudreve/Cloudreve)

vue 项目参照[V01 Vue.js 实战教程 - 基础篇](https://learnku.com/courses/vuejs-essential/2.6)

使用gin+gorm+statik+vue全家桶

前端目录为asset

运行教程
```
# 项目主目录
cd ../

# 安装 statik, 用于嵌入静态资源
go get github.com/rakyll/statik

# 开始嵌入
statik -src=assets/dist/  -include=*.html,*.js,*.json,*.css,*.png,*.svg,*.ico,*.eot,*.woff2,*.ttf,*.woff -f

go build main.go

chmod +x ./main

./main
```
默认使用sqlite数据库，切换到mysql可以修改conf.ini文件
