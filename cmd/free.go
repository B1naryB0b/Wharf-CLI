package cmd

import (
	"fmt"
	"strconv"

	"Wharf-CLI/internal/ports"

	"github.com/spf13/cobra"
)

var freeCmd = &cobra.Command{
	Use:     "free [count]",
	Aliases: []string{"f"},
	Short:   "Find the next available ports.",
	Long:    "Find the next available ports. Specify how many free ports you want to find.",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		count := 1
		err := error(nil)
		if len(args) > 0 {
			count, err = strconv.Atoi(args[0])
		}

		if err != nil {
			return err
		}

		return Free(count)
	},
}

func Free(count int) error {
	manager := ports.NewPlatformManager()

	if count < 0 {
		return fmt.Errorf("count must be greater than 0")
	}

	fmt.Printf("Finding free ports...\n")
	result, err := manager.GetNextFreePort(count)

	if result != "" {
		fmt.Println(result)
	}

	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(freeCmd)
}
