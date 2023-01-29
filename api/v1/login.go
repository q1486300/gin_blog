package v1

import (
	"gin_blog/middleware"
	"gin_blog/model"
	"gin_blog/utils/err_msg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Login 後台登入
func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)

	user, code := model.CheckLogin(data.UserName, data.Password)

	if code == err_msg.SUCCESS {
		setToken(c, user)
	} else {
		c.JSON(http.StatusOK, model.LoginResponse{
			Status:  code,
			Data:    user.UserName,
			ID:      user.ID,
			Message: err_msg.GetErrMsg(code),
		})
	}
}

// token 產生函數
func setToken(c *gin.Context, user model.User) {
	j := middleware.NewJWT()
	claims := middleware.MyClaims{
		UserName: user.UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Add(7 * time.Hour * 24).Unix(),
			Issuer:    "GinBlog",
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		code := err_msg.ERROR
		c.JSON(http.StatusOK, model.LoginResponse{
			Status:  code,
			Message: err_msg.GetErrMsg(code),
			Token:   token,
		})
	}

	code := err_msg.SUCCESS
	c.JSON(http.StatusOK, model.LoginResponse{
		Status:  code,
		Data:    user.UserName,
		ID:      user.ID,
		Message: err_msg.GetErrMsg(code),
		Token:   token,
	})
}
