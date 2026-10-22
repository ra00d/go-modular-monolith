package cli

import (
	"modular-monolith/lib/app"
	"modular-monolith/lib/logger"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var (
	appContainer *app.App
	serverCmd    = &cobra.Command{
		Use:   "start",
		Short: "Start the server",
		Run: func(cmd *cobra.Command, args []string) {
			if appContainer == nil {
				logger.Error("Please start the server first")
			}
			appContainer.Run()
		},
	}
)

func StartServer(mods ...fx.Option) {
	appContainer = app.NewApp(mods...)
	rootCmd.SetArgs([]string{"start"})
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
