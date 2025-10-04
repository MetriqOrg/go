package cmd

import (
	"github.com/spf13/cobra"
	orbitr "github.com/metriqorg/go/services/orbitr/internal"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run orbitr server",
	Long:  "serve initializes then starts the orbitr HTTP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := orbitr.NewAppFromFlags(config, flags)
		if err != nil {
			return err
		}
		return app.Serve()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
