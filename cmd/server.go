package cmd

import (
	"context"
	"fmt"
	"github.com/LynnWonder/gin_prac/api"
	"github.com/LynnWonder/gin_prac/pkg/common"
	"github.com/LynnWonder/gin_prac/pkg/db"
	_ "github.com/LynnWonder/gin_prac/pkg/db/model"
	"github.com/LynnWonder/gin_prac/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return viper.BindPFlags(cmd.LocalFlags())
	},
	Run: decoRunFunc(runServer),
}
var graced bool

type cmdFunc func(*cobra.Command, []string) error

func decoRunFunc(f cmdFunc) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		err := f(cmd, args)
		if err != nil {
			log.Fatalf("run cmd failed: %+v \n", err)
			os.Exit(1)
		}
	}
}

func runServer(cmd *cobra.Command, args []string) error {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// init mysql + autoMigrate
	db.Init(&common.AppConfig.DB, cmd)
	ctx, serverCancel := context.WithCancel(context.Background())
	go runHttpServe(cmd, ctx, common.AppConfig.Server.Port)
	gracechan := make(chan os.Signal, 1)
	signal.Notify(gracechan, syscall.SIGUSR1, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	exit(gracechan, serverCancel, serverCancel)
	return nil
}

func runHttpServe(cmd *cobra.Command, ctx context.Context, port int) {
	if !common.AppConfig.Server.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}
	router := api.NewRouter(cmd)
	defer func() {
		logger.Info(fmt.Sprintf("listening on 0.0.0.0:%d", port))
		_ = router.Run(fmt.Sprintf(":%d", port))
	}()
}

// TODO graceful shutdown
func exit(ch <-chan os.Signal, serverCancel, graceCancel context.CancelFunc) {
	sig := <-ch
	if sig == syscall.SIGUSR1 {
		logger.Info("Graceful quit signal received")
	} else {
		logger.Info("normal quit signal received")
	}

	if graced {
		return
	}

	// first close hc port
	graceCancel()

	// 4 seconds is safe to exit
	time.Sleep(4 * time.Second)

	// if graceful exit, wait another 10 seconds, give enough time to check us unhealthy
	if sig == syscall.SIGUSR1 {
		graced = true
		time.Sleep(10 * time.Second)
	}

	// close server port
	serverCancel()
	time.Sleep(4 * time.Second)
}
