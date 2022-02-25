package userRouter

import "github.com/gin-gonic/gin"

func UserAuthRouter(g * gin.RouterGroup){
	g.GET("/login",func(c *gin.Context) {
		c.JSON(200, gin.H{"msg":"登录成功"})
	})
}