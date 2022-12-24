package errors

import (
	"fmt"
	"gorm.io/gorm"

	"github.com/LynnWonder/gin_prac/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
)

type ICustomHTTPError interface {
	error
	Code() int
	Message() string
	HTTPCode() int
}

func CustomErrorHTTPHandler(err error) ICustomHTTPError {
	var e ICustomHTTPError
	switch err.(type) {
	case *BizHTTPError:
		e = err.(ICustomHTTPError)
	case *ServerHTTPError:
		e = err.(ICustomHTTPError)
	case *ClientHTTPError:
		e = err.(ICustomHTTPError)
	case *gin.Error:
		return e
	case *validator.InvalidValidationError, *validator.ValidationErrors, validator.ValidationErrors:
		ex := *ClientParamValueInvalidHTTPError
		ex.SetMsg(fmt.Sprintf("%s: %+v", ex.msg, err))
		e = &ex
		logger.Error(fmt.Sprintf("[CustomErrorHTTPHandler] validator error: %+v", err))
	case *mysql.MySQLError:
		mErr := err.(*mysql.MySQLError)
		if mErr.Number == 1062 || mErr.Number == 1069 {
			e = NewBizError("数据唯一性错误")
		} else {
			e = ServerDBHTTPError
			logger.Error(fmt.Sprintf("[CustomErrorHTTPHandler] mysql error, err: %+v", err))
		}
	default:
		if err == gorm.ErrRecordNotFound {
			e = NewBizError("数据不存在")
		} else {
			ex := *ServerUnknownHTTPError
			ex.SetMsg(fmt.Sprintf("%s: %+v", ex.msg, err))
			e = &ex
			logger.Error(fmt.Sprintf("[CustomErrorHTTPHandler] not match err type, err: %+v", err))
		}
	}
	return e
	//c.JSON(e.HTTPCode(), gin.H{"code": e.Code(), "msg": e.Message()})
}
