package api

import (
	"fmt"
	"net/http"

	"github.com/LynnWonder/gin_prac/pkg/common"
	"github.com/LynnWonder/gin_prac/pkg/common/errors"
	"github.com/LynnWonder/gin_prac/pkg/logger"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"github.com/LynnWonder/gin_prac/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	Name    string
	Method  string
	Path    string
	Handler func(ctx *gin.Context) (interface{}, error)
}

var apiV1Routes = []*Router{
	{
		Name:    "Get health check",
		Method:  "GET",
		Path:    "/health",
		Handler: GetHealthCheck,
	},
}

func GetHealthCheck(ctx *gin.Context) (interface{}, error) {
	return "hello world", nil
}

// NewRouter tip: 规范化输入和输出
func NewRouter(cmd *cobra.Command) *gin.Engine {
	gin.ForceConsoleColor()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(location.Default())

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	apiV1 := r.Group(basePath)
	for _, route := range apiV1Routes {
		route := route
		apiV1.Handle(route.Method, route.Path, func(ctx *gin.Context) {
			defer func() {
				if p := recover(); p != nil {
					ctx.JSON(http.StatusInternalServerError, p)
					panic(p)
				}
			}()

			// 从请求参数中获取 RequestId，如果 RequestId 为空，则自己生成
			requestId := getRequestId(ctx)
			data, err := route.Handler(ctx)
			if err != nil {
				logger.Error(fmt.Sprintf("Name: %s, Path: %s, cmerr %+v, data: %+v", route.Name, route.Path, err, data))
				// 规范 error
				customErr := errors.CustomErrorHTTPHandler(err)
				rsp := getErrRsp(requestId, customErr)
				ctx.JSON(customErr.HTTPCode(), rsp)
				return
			}
			if data != nil {
				rsp := getSuccessRsp(requestId, data)
				ctx.JSON(http.StatusOK, rsp)
				return
			}
		})
	}

	return r
}

func getRequestId(ctx *gin.Context) string {
	// 从请求参数中获取RequestId，如果 RequestId 为空，则自己生成
	requestId := ctx.Request.Header.Get("Gin-Request-Id")
	if len(requestId) == 0 {
		u4 := uuid.New()
		requestId = u4.String()
		ctx.Request.Header.Set("Gin-Request-Id", requestId)
	}
	return requestId
}

func getErrRsp(requestId string, err errors.ICustomHTTPError) *common.ErrorResponseMetadata {

	rsp := &common.ErrorResponseMetadata{
		BaseResponseMetadata: common.BaseResponseMetadata{RequestId: requestId},
		Error:                common.Error{Code: err.Code(), Message: err.Message()},
	}

	return rsp
}

func getSuccessRsp(requestId string, data interface{}) *common.ResponseMetadata {
	rsp := &common.ResponseMetadata{
		ResponseMetadata: &common.BaseResponseMetadata{RequestId: requestId},
		Result:           data,
	}
	return rsp
}
