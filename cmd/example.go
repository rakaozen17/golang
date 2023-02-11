package cmd

import (
	"fmt"
	app "restproject/app"

	"github.com/spf13/cobra"
)

func ExampleCommand(application *app.App) *cobra.Command {
	return &cobra.Command{
		Use:   "print [text to print]",
		Short: "Prints text to the console",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
	}
}
