package main

import (
	"github.com/LynnWonder/gin_prac/biz/handler/person"
	"github.com/LynnWonder/gin_prac/biz/handler/rpc"
	"github.com/gin-gonic/gin"
)

func register(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		personRouter := v1.Group("/person")
		{
			personRouter.GET("", person.ListAll)
			personRouter.POST("", person.CreatePerson)
			personRouter.PUT("/:personId", person.UpdatePerson)
			personRouter.DELETE("/:personId", person.DeletePerson)
		}
		rpcRouter := v1.Group("/rpc")
		{
			rpcRouter.POST("", rpc.Echo)
		}
	}
}
