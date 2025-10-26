package cmd

import (
	"fmt"

	"Wharf-CLI/internal/ports"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list [port]",
	Aliases: []string{"l"},
	Short:   "List all network ports or a specific port",
	Long:    "List all network ports currently running on your machine, or specify a port number to check",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return List(args[0])
	},
}

func List(port string) error {
	manager := ports.NewPlatformManager()

	if len(port) == 0 {
		fmt.Println("Listing all network ports")
		result, err := manager.GetAllPorts()
		if err != nil {
			return fmt.Errorf("failed to list ports: %w", err)
		}
		if result != "" {
			fmt.Println(result)
		} else {
			fmt.Println("No ports found.")
		}
	} else {
		fmt.Printf("Checking port %s\n", port)
		result, err := manager.GetPort(port)
		if err != nil {
			return fmt.Errorf("failed to get port %s: %w", port, err)
		}
		if result != "" {
			fmt.Println(result)
		} else {
			fmt.Printf("No process found on port %s\n", port)
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
}
