package routes

import (
	"github.com/gin-gonic/gin"

	v1 "goblog/api/v1"
	"goblog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	routerV1 := r.Group("api/v1")
	{
		// User模块的路由接口
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.GetUsers)
		routerV1.PUT("user/:id", v1.UpdateUser)
		routerV1.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口

		// 标签模块的路由接口

		// 文章模块的路由接口

		// 评论模块的路由接口

	}
	r.Run(utils.HttpPort)
}
