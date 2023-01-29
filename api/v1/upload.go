package v1

import (
	"fmt"
	"gin_blog/model"
	"gin_blog/utils/err_msg"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

func Upload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		code := err_msg.ERROR_UPLOAD_FILE
		c.JSON(http.StatusOK, model.Response{
			Status:  code,
			Message: err_msg.GetErrMsg(code),
		})
		return
	}

	var fileNames []string
	files := form.File["uploads[]"]
	for _, file := range files {
		fileName := strings.Split(file.Filename, ".")[0]
		fileExt := filepath.Ext(file.Filename)
		fileName = fmt.Sprintf("%s_%d%s", fileName, time.Now().UnixNano(), fileExt)

		c.SaveUploadedFile(file, fmt.Sprintf("%s%s", "./uploads/", fileName))
		fileNames = append(fileNames, fileName)
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  err_msg.SUCCESS,
		Data:    fileNames,
		Message: fmt.Sprintf("成功上傳 %d 個文件", len(fileNames)),
	})
}
