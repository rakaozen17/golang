package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func exampleCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "print [text to print]",
		Short: "Prints text to the console",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
	}
}
