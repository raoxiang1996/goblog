package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"goblog/model"
	"goblog/utils/errmsg"
)

// 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	if err := c.ShouldBindJSON(&data); err != nil {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, http.StatusBadRequest,
			errmsg.GetErrMsg(errmsg.PARSEBODYFAIL))
		c.JSON(http.StatusBadRequest, error)
		return
	}

	code := model.CheckCategory(data)
	if code == errmsg.SUCCESS {
		model.CreateCategory(&data)
	} else {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCategory(id)
	if code != errmsg.SUCCESS {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个分类

// 查询分类列表
func GetCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetCategory(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 修改分类
func UpdateCategory(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code := model.CheckCategory(data)
	if code != errmsg.SUCCESS {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, code,
			errmsg.GetErrMsg(code))
		c.JSON(http.StatusBadRequest, error)
		return
	}
	model.UpdateCategory(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
