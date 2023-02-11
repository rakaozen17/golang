package cmd

import (
	app "restproject/app"
	rest "restproject/rest"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func RestCommand(application *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "rest",
		Short: "Run a rest server",
		RunE: func(cmd *cobra.Command, args []string) error {
			logrus.WithField("component", "rest").Info("Running a rest server")
			rest.Serve(application)
			return nil
		},
	}
}
