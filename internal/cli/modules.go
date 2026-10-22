package cli

import (
	"fmt"

	"modular-monolith/internal/modules"

	"github.com/spf13/cobra"
)

var modulesCmd = &cobra.Command{
	Use:   "modules",
	Short: "List all registered modules",
	Run: func(cmd *cobra.Command, args []string) {
		for _, m := range modules.Meta() {
			fmt.Println(m)
			// fmt.Printf("%-15s %s\n", m.Name, m.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(modulesCmd)
}
