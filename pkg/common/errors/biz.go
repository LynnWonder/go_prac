package errors

import "net/http"

type BizHTTPError struct {
	code     int
	httpCode int
	msg      string
}

func (b *BizHTTPError) Error() string {
	return b.msg
}

func (b *BizHTTPError) Code() int {
	return b.code
}

func (b *BizHTTPError) Message() string {
	return b.msg
}

func (b *BizHTTPError) HTTPCode() int {
	return b.httpCode
}

// common biz error
func NewBizError(msg string) *BizHTTPError {
	return &BizHTTPError{
		msg:      msg,
		code:     30000,
		httpCode: http.StatusOK,
	}
}

func NewBizCloudenvError(msg string) *BizHTTPError {
	return &BizHTTPError{
		code:     30004,
		httpCode: http.StatusOK,
		msg:      msg,
	}
}

func NewBizComputeError(msg string) *BizHTTPError {
	return &BizHTTPError{
		msg:      msg,
		code:     30007,
		httpCode: http.StatusOK,
	}
}
