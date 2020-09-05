package routes

import (
	"github.com/gin-gonic/gin"
	"goblog/utils"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	router := r.Group("api/v1")
	router.GET("hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	},
	)
	r.Run(utils.HttpPort)
}
