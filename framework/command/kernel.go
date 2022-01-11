package command

import "github.com/gohade/hade/framework/cobra"

func AddKernelCommands(root *cobra.Command) {
	root.AddCommand(DemoCommand)

	root.AddCommand(initAppCommand())
}