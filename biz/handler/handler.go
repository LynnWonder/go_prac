package handler

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var trans ut.Translator
var ValiObj *validator.Validate

// Parse query with Application/json
func QueryValidate(r interface{}, c *gin.Context) *Custom_error {
	if err := c.Bind(r); err != nil {
		//  Gin’s validator will return validator.ValidationErrors if a validation error occurs,
		//  basically it will just send back the error it encountered.
		errs, ok := err.(validator.ValidationErrors)
		rawError := ErrHandlerQueryFail
		// so we do a type assertion here
		if !ok {
			rawError.Msg = err.Error()
			return rawError
		}
		rawError.Msg = errs.Translate(trans)
		return rawError
	}
	return nil
}

// Parse body with Application/json
func JSONValidate(r interface{}, c *gin.Context) *Custom_error {
	if err := c.BindJSON(r); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		rawError := ErrHandlerBodyFail
		if !ok {
			rawError.Msg = err.Error()
			return rawError
		}
		rawError.Msg = errs.Translate(trans)
		return rawError
	}
	return nil
}

// Parse params
func ParamsValidate(r interface{}, c *gin.Context) *Custom_error {
	if err := c.ShouldBindUri(r); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		rawError := ErrHandlerParamsFail
		if !ok {
			rawError.Msg = err.Error()
			return rawError
		}
		rawError.Msg = errs.Translate(trans)
		return rawError
	}
	return nil
}
