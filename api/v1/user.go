package v1

import (
	"fmt"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

// 查询用户是否存在
func UserIsExist() {

}

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Printf(err.Error())
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, http.StatusBadRequest,
			errmsg.GetErrMsg(errmsg.PARSEBODYFAIL))
		c.JSON(http.StatusBadRequest, error)
		return
	}

	code := model.CheckUser(data)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	} else {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, http.StatusBadRequest,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)
	if code != errmsg.SUCCESS {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, http.StatusBadRequest,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个用户

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 修改用户
func UpdateUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code := model.CheckUser(data)
	if code != errmsg.SUCCESS {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, http.StatusBadRequest,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}
	model.UpdateUser(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}
