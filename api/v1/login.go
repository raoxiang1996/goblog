package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goblog/middleware"
	"goblog/model"
	"goblog/utils/errmsg"
)

func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	code := model.CheckLogin(data.Username, data.Password)
	if code != errmsg.SUCCESS {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}
	token, code := middleware.SetToken(data.Username)
	if code != errmsg.SUCCESS {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
