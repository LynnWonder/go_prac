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

//ginZap "github.com/LynnWonder/gin_prac/biz/middleware/zap"
"github.com/gin-gonic/gin"
//"go.uber.org/zap"
)

func main() {
	// mode
	if !config.AppConfig.Server.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	defer func() {
		serve := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
		fmt.Printf("running service at %s\n", serve)
		_ = router.Run(serve)
	}()

	// Disabled TrustedProxies feature
	_ = router.SetTrustedProxies(nil)

	//// log
	//router.Use(ginZap.Ginzap(zap.L(), time.RFC3339, true))
	//router.Use(ginZap.RecoveryWithZap(zap.L(), true))

	// auth
	// router.Use(middleware.TokenAuthMiddleware())

	register(router)

}
