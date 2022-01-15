package command

import (
	"fmt"

	"github.com/gohade/hade/framework/cobra"
	"github.com/gohade/hade/framework/contract"
	"github.com/kr/pretty"
)

func initConfigCommand() *cobra.Command {
	configCommand.AddCommand(configGetCommand)
	return configCommand
}

var configCommand=&cobra.Command{
	Use: "config",
	Short: "获取配置的相关信息",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args)==0 {
			cmd.Help()
		}
		return nil
	},
}

var configGetCommand =&cobra.Command{
	Use: "get",
	Short: "获取某个配置的信息",
	RunE: func(cmd *cobra.Command, args []string) error {
		contianer:=cmd.GetContainer()
		configService:=contianer.MustMake(contract.ConfigKey).(contract.Config)
		if len(args) !=1{
			fmt.Println("参数错误")
			return nil
		}
		configPath:=args[0]
		val:=configService.Get(configPath)
		if val==nil {
			fmt.Println("配置路径: ",configPath,"不存在")
			return nil
		}
		fmt.Printf("%# v\n",pretty.Formatter(val))
		return nil
	},
}