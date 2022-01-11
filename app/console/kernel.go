package console

import (
	"github.com/gohade/hade/app/console/command/demo"
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/cobra"
	"github.com/gohade/hade/framework/command"
)

//初始化根Command并运行
func RunCommand(container framework.Container) error {
	//根Command
	var rootCmd =&cobra.Command{
		Use: "hade",
		Short: "hade 命令",
		Long: "hade 框架提供的命令行工具，使用这个命令行工具能很方便执行框架自带命令，也能够很方便编写业务命令",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		//不需要出现cobra默认的completion子命令
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	//为根Command设置服务容器
	rootCmd.SetContainer(container)
	//绑定框架的命令
	command.AddKernelCommands(rootCmd)
	//绑定业务的命令
	AddAppCommand(rootCmd)

	//执行RootCommand
	return rootCmd.Execute()
}

//绑定业务的命令
func AddAppCommand(rootCmd *cobra.Command){
	rootCmd.AddCommand(demo.InitFoo())
}