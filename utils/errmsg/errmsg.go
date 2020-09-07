package errmsg

import "errors"

const (
	SUCCESS = 200
	ERROR   = 500
	// code 1000 用户模块的错误
	ERROR_UERNAME_USED     = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	// code 2000 文章模块错误

	// code 3000 分类模块错误

	// code 4000 评论模块错误

)

var codemsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	ERROR_UERNAME_USED:     "用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_NOT_EXIST:  "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式不正确",
}

func GetErrMsg(code int) (string, error) {
	errmsg, ok := codemsg[code]
	if ok {
		return "", errors.New("get error message fail")
	}
	return errmsg, nil
}
