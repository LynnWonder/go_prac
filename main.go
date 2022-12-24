//package main
//
//import (
//	"fmt"
//	"github.com/LynnWonder/gin_prac/biz/config"
//	"github.com/LynnWonder/gin_prac/router"
//	"github.com/gin-gonic/gin"
//	"go.uber.org/zap"
//	"time"
//)
//func main() {
//	r :=router.InitRouter()
//	err := r.Run(":8080")
//	if err != nil {
//		fmt.Println("服务器启动失败！")
//	}
//}

package main

import (
	"github.com/LynnWonder/gin_prac/cmd"
)

func main() {
	cmd.Execute()
}
