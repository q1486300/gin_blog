package err_msg

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000... 用戶模組的錯誤
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_TYPE_WRONG = 1006
	ERROR_USER_NO_RIGHT    = 1007

	// code = 2000... 文章模組的錯誤
	ERROR_ARTICLE_NOT_EXIST = 2001

	// code = 3000... 分類模組的錯誤
	ERROR_CATENAME_USED = 3001

	// code = 5000... 文件相關的錯誤
	ERROR_UPLOAD_FILE = 5001
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用戶名已存在",
	ERROR_PASSWORD_WRONG:   "密碼錯誤",
	ERROR_USER_NOT_EXIST:   "用戶不存在",
	ERROR_TOKEN_NOT_EXIST:  "Token不存在",
	ERROR_TOKEN_RUNTIME:    "Token授權已過期，請重新登入",
	ERROR_TOKEN_TYPE_WRONG: "Token格式錯誤",
	ERROR_USER_NO_RIGHT:    "此用戶無權限",

	ERROR_ARTICLE_NOT_EXIST: "文章不存在",

	ERROR_CATENAME_USED: "分類已存在",

	ERROR_UPLOAD_FILE: "文件上傳失敗",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
