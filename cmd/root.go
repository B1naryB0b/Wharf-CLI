package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wharf",
	Short: "Wharf is a CLI tool to help you manage network ports on any platform.",
	Long:  "Wharf is a CLI tool to help you manage network ports on any platform. You can manage ports in groups or individually.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
