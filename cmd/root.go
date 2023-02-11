package cmd

import (
	"os"
	app "restproject/app"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Execute() error {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	var root = &cobra.Command{}
	application := app.InitApp()
	root.AddCommand(
		exampleCommand(application),
	)

	return root.Execute()
}
