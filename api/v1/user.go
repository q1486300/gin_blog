package v1

import (
	"gin_blog/model"
	"gin_blog/utils/err_msg"
	"gin_blog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加用戶
func AddUser(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)

	msg, validCode := validator.Validate(&data)
	if validCode != err_msg.SUCCESS {
		c.JSON(http.StatusOK, model.Response{
			Status:  validCode,
			Message: msg,
		})
		c.Abort()
		return
	}

	code := model.CheckUser(data.UserName)
	if code == err_msg.SUCCESS {
		code = model.CreateUser(&data)
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  code,
		Message: err_msg.GetErrMsg(code),
	})
}

// 查詢某一位用戶

// 查詢用戶列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum <= 0 {
		pageNum = 1
	}

	data, total := model.GetUsers(username, pageSize, pageNum)
	code := err_msg.SUCCESS

	c.JSON(http.StatusOK, model.Response{
		Status:  code,
		Data:    data,
		Total:   total,
		Message: err_msg.GetErrMsg(code),
	})
}

// 編輯用戶
func EditUser(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)

	msg, validCode := validator.Validate(&data, "Password")
	if validCode != err_msg.SUCCESS {
		c.JSON(http.StatusOK, model.Response{
			Status:  validCode,
			Message: msg,
		})
		c.Abort()
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	code := model.CheckUser(data.UserName)
	if code == err_msg.SUCCESS {
		code = model.EditUser(id, &data)
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  code,
		Message: err_msg.GetErrMsg(code),
	})
}

// 刪除用戶
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteUser(id)

	c.JSON(http.StatusOK, model.Response{
		Status:  code,
		Message: err_msg.GetErrMsg(code),
	})
}
