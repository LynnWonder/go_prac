package rpc

import (
	"context"
	"github.com/LynnWonder/gin_prac/biz/common"
	"github.com/LynnWonder/gin_prac/biz/handler"
	"github.com/LynnWonder/gin_prac/rpc/kitex_gen/api"
	"github.com/LynnWonder/gin_prac/rpc/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)


// 绑定 JSON
type EchoRequest struct {
	Message       string `json:"message" binding:"required,lte=50"`
}

var log = common.SugaredLogger

func Echo(c *gin.Context) {
	var r EchoRequest

	// 校验传进来的参数
	if err := handler.JSONValidate(&r, c); err != nil {
		c.JSON(err.HTTPCode, err.ToMap())
		return
	}

	// 下述逻辑是编写了一个客户端调用我们在 rpc 文件夹已经运行起来的服务端
	contextE, err := echo.NewClient("example", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Errorf("echo NewClient error: %v", err)
	}

	// 表示从一个变量身上取址，* 则是根据地址取值，它也表示一个指针类型
	req := &api.Request{Message: r.Message}
	resp, err := contextE.Echo(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Errorf("request Echo error: %v", err)
	}
	c.JSON(http.StatusOK, resp)
}