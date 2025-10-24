package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list [port]",
	Aliases: []string{"l"},
	Short:   "List all network ports or a specific port",
	Long:    "List all network ports currently running on your machine, or specify a port number to check",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Listing all network ports")
			result := ListPorts("")
			if result != "" {
				fmt.Println(result)
			} else {
				fmt.Println("No ports found.")
			}
		} else {
			port := args[0]
			fmt.Printf("Checking port %s\n", port)
			result := GetSpecifiedPort(port)
			if result != "" {
				fmt.Println(result)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
