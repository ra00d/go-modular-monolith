/*
	Package cmd

*
Copyright © 2026 Ra'ad A. Noman
*/
package cmd

import (
	"modular-monolith/tools/cli/internal/db"

	"github.com/spf13/cobra"
)

var (
	migrationPath      string
	createMigrationCmd = &cobra.Command{
		Use:   "migration:create [name]",
		Short: "Create new migration file",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return db.CreateNewMigration(args[0], migrationPath)
		},
	}
	migrateCmd = &cobra.Command{
		Use:   "db:migrate",
		Short: "migrate the database and generate jet code",
		RunE: func(cmd *cobra.Command, args []string) error {
			return db.MigrateDB(migrationPath)
		},
	}
)

func init() {
	rootCmd.AddCommand(createMigrationCmd, migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	rootCmd.PersistentFlags().
		StringVarP(&migrationPath, "module", "m", "", "Module directory for the migration, relative to the internal folder")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
