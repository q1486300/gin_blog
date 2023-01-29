package middleware

import (
	"errors"
	"gin_blog/model"
	"gin_blog/utils"
	"gin_blog/utils/err_msg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(utils.JwtKey),
	}
}

type MyClaims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

// 定義錯誤
var (
	TokenExpired     = errors.New("token已過期，請重新登入")
	TokenNotValidYet = errors.New("token無效，請重新登入")
	TokenMalformed   = errors.New("token不正確，請重新登入")
	TokenInvalid     = errors.New("這不是一個token，請重新登入")
)

// CreateToken 產生token
func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// ParserToken 解析token
func (j *JWT) ParserToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, TokenInvalid
}

// JwtToken jwt中間件
func JwtToken(c *gin.Context) {
	tokenHeader := c.GetHeader("Authorization")
	if tokenHeader == "" {
		code := err_msg.ERROR_TOKEN_NOT_EXIST
		c.JSON(http.StatusOK, model.Response{
			Status:  code,
			Message: err_msg.GetErrMsg(code),
		})
		c.Abort()
		return
	}

	checkToken := strings.Split(tokenHeader, " ")
	if len(checkToken) == 0 {
		code := err_msg.ERROR_TOKEN_NOT_EXIST
		c.JSON(http.StatusOK, model.Response{
			Status:  code,
			Message: err_msg.GetErrMsg(code),
		})
		c.Abort()
		return
	}

	if len(checkToken) != 2 || checkToken[0] != "Bearer" {
		code := err_msg.ERROR_TOKEN_TYPE_WRONG
		c.JSON(http.StatusOK, model.Response{
			Status:  code,
			Message: err_msg.GetErrMsg(code),
		})
		c.Abort()
		return
	}

	j := NewJWT()
	claims, err := j.ParserToken(checkToken[1])
	if err != nil {
		if err == TokenExpired {
			code := err_msg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, model.Response{
				Status:  code,
				Message: err_msg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		// 其他錯誤
		c.JSON(http.StatusOK, model.Response{
			Status:  err_msg.ERROR,
			Message: err.Error(),
		})
		c.Abort()
		return
	}

	c.Set("username", claims.UserName)
	c.Next()
}
