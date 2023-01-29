package v1

import (
	"gin_blog/model"
	"gin_blog/utils/err_msg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 查詢分類是否存在

// 添加分類
func AddCategory(c *gin.Context) {
	var data model.Category
	c.ShouldBindJSON(&data)

	code := model.CheckCategory(data.Name)
	if code == err_msg.SUCCESS {
		code = model.CreateCategory(&data)
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  code,
		Data:    data,
		Message: err_msg.GetErrMsg(code),
	})
}

// 查詢某一個分類下的文章

// 查詢分類列表
func GetCategorys(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum <= 0 {
		pageNum = 1
	}

	data, total := model.GetCategory(pageSize, pageNum)
	code := err_msg.SUCCESS

	c.JSON(http.StatusOK, model.Response{
		Status:  code,
		Data:    data,
		Total:   total,
		Message: err_msg.GetErrMsg(code),
	})
}

// 編輯分類
func EditCategory(c *gin.Context) {
	var data model.Category
	c.ShouldBindJSON(&data)

	id, _ := strconv.Atoi(c.Param("id"))

	code := model.CheckCategory(data.Name)
	if code == err_msg.SUCCESS {
		code = model.EditCategory(id, &data)
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  code,
		Message: err_msg.GetErrMsg(code),
	})

}

// 刪除分類
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteCategory(id)

	c.JSON(http.StatusOK, model.Response{
		Status:  code,
		Message: err_msg.GetErrMsg(code),
	})
}
