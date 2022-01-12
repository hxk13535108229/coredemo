package command

import (
	"fmt"

	"github.com/gohade/hade/framework/cobra"
	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/util"
)


//获取env相关的命令
func initEnvCommand() *cobra.Command {
	envCommand.AddCommand(envListCommand)
	return envCommand
}

//获取当前的App环境
var envCommand = &cobra.Command{
	Use: "env",
	Short: "获取当前的App环境",
	Run: func(cmd *cobra.Command, args []string) {
		//获取env环境
		container := cmd.GetContainer()
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		//打印
		fmt.Println("environment: ",envService.AppEnv())
	},
}

//获取所有的App环境变量
var envListCommand = &cobra.Command{
	Use: "list",
	Short: "获取所有的App环境变量",
	Run: func(cmd *cobra.Command, args []string) {
			//获取env环境
			container := cmd.GetContainer()
			envService := container.MustMake(contract.EnvKey).(contract.Env)
			envs :=envService.All()
			outs := [][]string {}
			for k,v := range envs{
				outs = append(outs, []string{k,v})
			}
			util.PrettyPrint(outs)
	},
}