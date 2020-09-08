package v1

import (
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	if err := c.ShouldBindJSON(&data); err != nil {
		error := errmsg.SetErrorResponse(c.Request.Method, c.Request.URL.Path, http.StatusBadRequest,
			errmsg.GetErrMsg(errmsg.PARSEBODYFAIL))
		c.JSON(http.StatusBadRequest, error)
		return
	}
	code := model.CreateArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除文章
func DeleteArticle(c *gin.Context) {
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

// 查询单个文章

// 查询文章列表
func GetArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetArticle(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 修改文章
func UpdateArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code := model.UpdateArticle(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
