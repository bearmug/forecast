package cmd

import (
	"fmt"
	"github.com/bearmug/forecast/config"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Run the interactive configuration setup",
	Long:  "The setup command allows you to create a new configuration or modify your existing one by running an interactive setup process.",
	Run: func(cmd *cobra.Command, args []string) {
		err := config.CreateDefaultConfig()
		if err != nil {
			fmt.Printf("Error during configuration: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
