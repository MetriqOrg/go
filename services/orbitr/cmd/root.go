package cmd

import (
	"fmt"
	stdLog "log"

	"github.com/spf13/cobra"
	orbitr "github.com/metriqorg/go/services/orbitr/internal"
)

var (
	config, flags = orbitr.Flags()

	RootCmd = &cobra.Command{
		Use:           "orbitr",
		Short:         "client-facing api server for the Lantah Network",
		SilenceErrors: true,
		SilenceUsage:  true,
		Long:          "Client-facing API server for the Lantah Network. It acts as the interface between Gravity and applications that want to access the Lantah Network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams and more.",
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := orbitr.NewAppFromFlags(config, flags)
			if err != nil {
				return err
			}
			return app.Serve()
		},
	}
)

// ErrUsage indicates we should print the usage string and exit with code 1
type ErrUsage struct {
	cmd *cobra.Command
}

func (e ErrUsage) Error() string {
	return e.cmd.UsageString()
}

// ErrExitCode Indicates we want to exit with a specific error code without printing an error.
type ErrExitCode int

func (e ErrExitCode) Error() string {
	return fmt.Sprintf("exit code: %d", e)
}

func init() {
	err := flags.Init(RootCmd)
	if err != nil {
		stdLog.Fatal(err.Error())
	}
}

func Execute() error {
	return RootCmd.Execute()
}
