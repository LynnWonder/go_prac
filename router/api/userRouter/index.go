package userRouter

import "github.com/gin-gonic/gin"

func UserInitRouter(r *gin.RouterGroup){
	userAuth := r.Group("/user-auth")
	UserAuthRouter(userAuth)
}