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
	"fmt"
	"github.com/LynnWonder/gin_prac/biz/config"
	_ "github.com/LynnWonder/gin_prac/biz/dal"
	ginZap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func main() {
	// mode
	if !config.AppConfig.Server.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	// gin 初始化
	router := gin.New()
	// Disabled TrustedProxies feature
	_ = router.SetTrustedProxies(nil)

	// log middleware
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	router.Use(ginZap.Ginzap(zap.L(), time.RFC3339, true))
	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(ginZap.RecoveryWithZap(zap.L(), true))

	// 注册路由
	register(router)
	// defer 函数在 return 之前执行
	defer func() {
		serve := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
		fmt.Printf("running service at %s\n", serve)
		_ = router.Run(serve)
	}()
}
