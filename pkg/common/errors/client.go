package errors

import "net/http"

type ClientHTTPError struct {
	httpCode int
	code     int
	msg      string
}

func (c *ClientHTTPError) Error() string {
	return c.msg
}

func (c *ClientHTTPError) Code() int {
	return c.code
}

func (c *ClientHTTPError) Message() string {
	return c.msg
}

func (c *ClientHTTPError) HTTPCode() int {
	return c.httpCode
}

func (c *ClientHTTPError) SetMsg(msg string) {
	c.msg = msg
}

var (
	ClientLoginRequiredHTTPError = &ClientHTTPError{
		code:     20001,
		msg:      "用户操作需登录",
		httpCode: http.StatusUnauthorized,
	}

	ClientParamValueInvalidHTTPError = &ClientHTTPError{
		code:     20002,
		msg:      "请求参数不合法或缺失",
		httpCode: http.StatusBadRequest,
	}

	ClientPermissionRequiredHTTPError = &ClientHTTPError{
		code:     20003,
		msg:      "用户操作需授权",
		httpCode: http.StatusForbidden,
	}
	ClientTokenInvalidError = &ClientHTTPError{
		code:     2004,
		msg:      "请求 token 无效",
		httpCode: http.StatusUnauthorized,
	}
	ClientTokenTimestampInvalidError = &ClientHTTPError{
		code:     2005,
		msg:      "请求 token 时间戳不合法",
		httpCode: http.StatusUnauthorized,
	}
	ClientTokenTimestampExpiredError = &ClientHTTPError{
		code:     2006,
		msg:      "请求 token 时间戳已过期",
		httpCode: http.StatusUnauthorized,
	}
)
