package api

import (
	"gin_prac/router/api/userRouter"
	"github.com/gin-gonic/gin"
)

// 创建路由组
func InitApi(r *gin.Engine) {
	api :=r.Group("/api")
	userRouter.UserInitRouter(api)
}