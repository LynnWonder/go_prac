package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	debug bool
)

var rootCmd = &cobra.Command{
	Use:   "prac",
	Short: "prac",
	Long:  "practice",
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug mode, output verbose output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("run cmd failed, err: %+v", err)
		os.Exit(1)
	}
}
