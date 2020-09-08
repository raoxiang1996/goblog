package errmsg

type Error struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

const (
	SUCCESS       = 200
	PARSEBODYFAIL = 412
	ERROR         = 500
	// code 1000 用户模块的错误
	ERROR_UERNAME_USED     = 1001
	ERROR_UERNAME_EMPTY    = 1002
	ERROR_PASSWORD_EMPTY   = 1003
	ERROR_PASSWORD_WRONG   = 1004
	ERROR_USER_NOT_EXIST   = 1005
	ERROR_TOKEN_NOT_EXIST  = 1006
	ERROR_TOKEN_RUNTIME    = 1007
	ERROR_TOKEN_WRONG      = 1008
	ERROR_TOKEN_TYPE_WRONG = 1009
	// code 2000 文章模块错误

	// code 3000 分类模块错误

	// code 4000 评论模块错误

)

var codemsg = map[int]string{
	SUCCESS:       "OK",
	PARSEBODYFAIL: "Parsing body failed",
	ERROR:         "Fail",

	ERROR_UERNAME_USED:     "用户名已存在",
	ERROR_UERNAME_EMPTY:    "用户名为空",
	ERROR_PASSWORD_EMPTY:   "密码为空",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_NOT_EXIST:  "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式不正确",
}

func GetErrMsg(code int) string {
	errmsg := codemsg[code]
	return errmsg
}

func SetErrorResponse(typ string, title string, status int, message string) Error {
	err := Error{
		Type:    typ,
		Title:   title,
		Status:  status,
		Message: message,
	}
	return err
}
