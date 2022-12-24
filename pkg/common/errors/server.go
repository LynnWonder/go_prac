package errors

import "net/http"

type ServerHTTPError struct {
	httpCode int
	code     int
	msg      string
}

func (s *ServerHTTPError) Error() string {
	return s.msg
}

func (s *ServerHTTPError) Code() int {
	return s.code
}

func (s *ServerHTTPError) Message() string {
	return s.msg
}

func (s *ServerHTTPError) HTTPCode() int {
	return s.httpCode
}

func (s *ServerHTTPError) SetMsg(msg string) {
	s.msg = msg
}

var (
	ServerUnknownHTTPError = &ServerHTTPError{
		code:     10001,
		msg:      "服务器未知错误",
		httpCode: http.StatusInternalServerError,
	}

	ServerDBHTTPError = &ServerHTTPError{
		code:     10002,
		msg:      "数据库错误",
		httpCode: http.StatusInternalServerError,
	}

	ServerCacheHTTPError = &ServerHTTPError{
		code:     10003,
		msg:      "缓存错误",
		httpCode: http.StatusInternalServerError,
	}
	ServerRpcHTTPError = &ServerHTTPError{
		code:     10004,
		msg:      "rpc 调用错误",
		httpCode: http.StatusInternalServerError,
	}
)
