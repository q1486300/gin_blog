package main

import (
	"gin_blog/model"
	"gin_blog/routes"
)

func main() {
	// 引用資料庫
	model.InitDb()

	routes.InitRouter()
}
