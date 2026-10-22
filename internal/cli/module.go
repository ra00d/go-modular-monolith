package cli

import (
	"modular-monolith/internal/cli/modulegen"

	"github.com/spf13/cobra"
)

var (
	moduleDir       string
	createModuleCmd = &cobra.Command{
		Use:     "module:create [name]",
		Aliases: []string{"module:generate", "module:new"},
		Short:   "Create a new module scaffold",
		Long: `Create a new module with api/rest, app, domain, infra, and migrations packages.

By default the module is created under internal/modules when that directory
exists, otherwise under internal. Use --path to choose another destination.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return modulegen.CreateModule(args[0], modulegen.Options{Dir: moduleDir})
		},
	}
)

func init() {
	createModuleCmd.Flags().
		StringVarP(&moduleDir, "path", "p", "", "Module destination relative to the project root (default: auto-detect)")

	rootCmd.AddCommand(createModuleCmd)
}
