package command

import "github.com/gohade/hade/framework/cobra"

func AddKernelCommands(root *cobra.Command) {
	//cron
	root.AddCommand(initCronCommand())

	//app
	root.AddCommand(initAppCommand())

	//env 
	root.AddCommand(initEnvCommand())

	//config
	root.AddCommand(initConfigCommand())
}