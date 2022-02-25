package main

import (
	"fmt"
	"gin_prac/router"
	)
func main() {
	//S :=gin.Default()
	//S.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"msg":"服务器启动成功"})
	//})
	r :=router.InitRouter()
	//err :=S.Run(":8080")
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("服务器启动失败！")
	}
}
