package routes

import (
	v1 "gin_blog/api/v1"
	"gin_blog/log"
	"gin_blog/middleware"
	"gin_blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)

	log.InitLog("logs", "GinBlog")

	r := gin.New()
	r.Use(middleware.Logger)
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.StaticFS("/uploads", http.Dir("./uploads"))

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken)
	{
		// 用戶模組的路由
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		// 分類模組的路由
		auth.POST("category", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)

		// 文章模組的路由
		auth.POST("article", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)

		// 上傳文件
		auth.POST("upload", v1.Upload)
	}

	router := r.Group("api/v1")
	{
		// 用戶模組的路由
		router.POST("user", v1.AddUser)
		router.GET("users", v1.GetUsers)

		// 分類模組的路由
		router.GET("categorys", v1.GetCategorys)

		// 文章模組的路由
		router.GET("articles", v1.GetArticles)
		router.GET("articles/:id", v1.GetArticleInfo)
		router.GET("articles/category/:cid", v1.GetCateArticle)

		// 登入控制模組的路由
		router.POST("login", v1.Login)
	}

	r.Run(utils.HttpPort)
}
