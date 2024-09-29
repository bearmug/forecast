package cmd

import (
	"fmt"
	"github.com/bearmug/forecast/config"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "forecast",
	Short: "A CLI tool for fetching weather forecasts",
	Long:  "Forecast is a CLI application written in Go that fetches weather forecasts for specified locations.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)
}
