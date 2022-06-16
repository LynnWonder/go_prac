package handler

import (
	"github.com/LynnWonder/gin_prac/biz/service/person"
	"github.com/gin-gonic/gin"
)

type Custom_error struct {
	Msg      interface{}
	Code     int
	HTTPCode int
	Extra    interface{}
}

var table = map[error]*Custom_error{
	person.ErrPersonNotFound:       ErrPersonNotFound,
}

func (ce *Custom_error) ToMap() map[string]interface{} {
	out := map[string]interface{}{"message": ce.Msg, "code": ce.Code}
	if ce.Extra != nil {
		out["extra"] = ce.Extra
	}
	return out
}

func ErrorsHandler(err error, c *gin.Context) {
	if ce := table[err]; table[err] != nil {
		c.JSON(ce.HTTPCode, map[string]interface{}{"message": ce.Msg, "code": ce.Code})
		return
	} else {
		c.JSON(ce.HTTPCode, map[string]interface{}{"message": ce.Msg, "code": ce.Code, "extra": err})
		return
	}
}

// Internal Error
var (
	UnknownError  = &Custom_error{Code: 10000, Msg: "Undefined err. Please contact with Admin.", HTTPCode: 500}
	GormError     = &Custom_error{Code: 10001, Msg: "Gorm error. Please contact with Admin", HTTPCode: 500}
	DBError       = &Custom_error{Code: 10002, Msg: "DB error. Please contact with Admin", HTTPCode: 500}
	InternalError = &Custom_error{Code: 10003, Msg: "Internal err. Please contact with Admin.", HTTPCode: 500}
)

// General Error
var (
	ErrHandlerQueryFail  = &Custom_error{Code: 20000, Msg: "Query is not satified with the need. Please check the doc or contact with dev team", HTTPCode: 400}
	ErrHandlerBodyFail   = &Custom_error{Code: 20001, Msg: "Body is not satified with the need. Please check the doc or contact with dev team", HTTPCode: 400}
	ErrHandlerParamsFail = &Custom_error{Code: 20002, Msg: "Params is not satified with the need. Please check the doc or contact with dev team", HTTPCode: 400}
)

// Person Error
var (
	ErrPersonNotFound    = &Custom_error{Code: 30001, Msg: "person not found", HTTPCode: 400}
)

