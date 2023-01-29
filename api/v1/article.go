package v1

import (
	"gin_blog/model"
	"gin_blog/utils/err_msg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	c.ShouldBindJSON(&data)

	code := model.CreateArticle(&data)

	c.JSON(http.StatusOK, model.Response{
		Status:  code,
		Data:    data,
		Message: err_msg.GetErrMsg(code),
	})
}

// 查詢某分類下的所有文章
func GetCateArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	cid, _ := strconv.Atoi(c.Param("cid"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum <= 0 {
		pageNum = 1
	}

	data, total := model.GetCateArticle(cid, pageSize, pageNum)
	code := err_msg.SUCCESS
	c.JSON(http.StatusOK, model.Response{
		Status:  code,
		Data:    data,
		Total:   total,
		Message: err_msg.GetErrMsg(code),
	})
}

// 查詢某一篇文章
func GetArticleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetArticleInfo(id)
	c.JSON(http.StatusOK, model.Response{
		Status:  code,
		Data:    data,
		Message: err_msg.GetErrMsg(code),
	})
}

// 查詢文章列表
func GetArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	title := c.Query("title")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum <= 0 {
		pageNum = 1
	}

	if len(title) == 0 {
		data, total := model.GetArticle(pageSize, pageNum)
		code := err_msg.SUCCESS
		c.JSON(http.StatusOK, model.Response{
			Status:  code,
			Data:    data,
			Total:   total,
			Message: err_msg.GetErrMsg(code),
		})
		return
	}

	data, total := model.SearchArticle(title, pageSize, pageNum)
	code := err_msg.SUCCESS
	c.JSON(http.StatusOK, model.Response{
		Status:  code,
		Data:    data,
		Total:   total,
		Message: err_msg.GetErrMsg(code),
	})
}

// 編輯文章
func EditArticle(c *gin.Context) {
	var data model.Article
	c.ShouldBindJSON(&data)

	id, _ := strconv.Atoi(c.Param("id"))

	code := model.EditArticle(id, &data)

	c.JSON(http.StatusOK, model.Response{
		Status:  code,
		Message: err_msg.GetErrMsg(code),
	})
}

// 刪除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteArticle(id)

	c.JSON(http.StatusOK, model.Response{
		Status:  code,
		Message: err_msg.GetErrMsg(code),
	})
}
