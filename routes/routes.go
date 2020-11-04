package routes

import (
	"github.com/gin-gonic/gin"

	v1 "goblog/api/v1"
	"goblog/middleware"
	"goblog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Log())
	r.Use(gin.Recovery()) // Recovery中间件可以让我们从崩溃中恢复
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// User模块的路由接口
		auth.PUT("user/:id", v1.UpdateUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.UpdateCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		// 标签模块的路由接口

		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.UpdateArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		// 评论模块的路由接口

	}
	router := r.Group("api/v1")
	{
		// User模块的路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		// 分类模块的路由接口
		router.GET("category", v1.GetCategory)
		// 标签模块的路由接口
		// 文章模块的路由接口
		router.GET("article", v1.GetArticle)
		router.GET("article/category", v1.GetCateArticle)
		router.GET("article/info/:id", v1.GetArticleInfo)
		router.POST("login", v1.Login)
		// 评论模块的路由接口
	}
	r.Run(utils.HttpPort)
}
