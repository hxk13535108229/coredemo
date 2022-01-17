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

	//build命令
	root.AddCommand(initBuildCommand())

	//npm build
	root.AddCommand(npmCommand)

	//go build
	root.AddCommand(goCommand)

	//dev
	root.AddCommand(initDevCommand())

	//provider
	root.AddCommand(initProviderCommand())

	//cmd
	root.AddCommand(initCmdCommand())
}