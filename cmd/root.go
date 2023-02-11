package cmd

import "github.com/spf13/cobra"

func Execute() error {
	var root = &cobra.Command{}
	root.AddCommand(
		exampleCommand(),
	)

	return root.Execute()
}
