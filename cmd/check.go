package cmd

import (
	"fmt"
	"strconv"

	"Wharf-CLI/internal/ports"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:     "check [port] [timeout]",
	Aliases: []string{"c"},
	Short:   "Check a specific port",
	Long:    "Check a specific port with a timeout",
	Args:    cobra.MaximumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		port := ""
		timeout := 1.0
		if len(args) > 0 {
			port = args[0]
		}
		if len(args) > 1 {
			timeout, _ = strconv.ParseFloat(args[1], 64)
		}

		return Check(port, timeout)
	},
}

func Check(port string, timeout float64) error {
	manager := ports.NewPlatformManager()

	if len(port) == 0 {
		return fmt.Errorf("port is required for checking")
	}

	fmt.Printf("Checking port %s...\n", port)
	result, err := manager.PingPort(port, timeout)

	if result != "" {
		fmt.Println(result)
	}

	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
