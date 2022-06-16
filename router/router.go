package router

import (
	"github.com/LynnWonder/gin_prac/router/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	// tip 这个中间件的作用是 程序如果出现 panic，那么中间件将会写入 500
	r.Use(gin.Recovery())
	// 新增路由
	api.InitApi(r)
	return r
}