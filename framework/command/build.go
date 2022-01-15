package command

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/gohade/hade/framework/cobra"
)

func initBuildCommand() *cobra.Command {
	buildCommand.AddCommand(buildSelfCommand)
	buildCommand.AddCommand(buildBackendCommand)
	buildCommand.AddCommand(buildFrontendCommand)
	buildCommand.AddCommand(buildAllCommand)
	return buildCommand
}

var buildCommand =&cobra.Command{
	Use: "build",
	Short: "编译相关命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args)==0 {
			cmd.Help()
		}
		return nil
	},
}

var buildSelfCommand=&cobra.Command{
	Use: "self",
	Short: "编译hade命令",
	RunE: func(c *cobra.Command, args []string) error {
		path,err:=exec.LookPath("go")
		if err!=nil {
			log.Fatalln("hade go: please install go in path first")
		}

		cmd :=exec.Command(path,"build","-o","hade","./")
		out,err:=cmd.CombinedOutput()
		if err!=nil {
			fmt.Println("go build error:")
			fmt.Println(string(out))
			fmt.Println("----------------")
			return err
		}
		fmt.Println("build success please run ./hade direct")
		return nil
	},
}

var buildBackendCommand=&cobra.Command{
	Use: "backend",
	Short: "使用go编译后端",
	RunE: func(cmd *cobra.Command, args []string) error {
		return buildSelfCommand.RunE(cmd,args)
	},
}

var buildFrontendCommand=&cobra.Command{
	Use: "frontend",
	Short: "使用npm编译前端",
	RunE: func(c *cobra.Command, args []string) error {
		path,err:=exec.LookPath("npm")
		if err!=nil {
			log.Fatalln("请安装npm在你的PATH路径下")
		}

		cmd := exec.Command(path,"run","build")
		out,err :=cmd.CombinedOutput()

		if err!=nil {
			fmt.Println("================ 前端编译失败 ================")
			fmt.Println(string(out))
			fmt.Println("================ 前端编译失败 ================")
			return err
		}
		//打印
		fmt.Println(string(out))
		fmt.Println("================ 前端编译成功 ================")
		return nil
	},
}

var buildAllCommand = &cobra.Command{
	Use: "all",
	Short: "同时编译前端和后端",
	RunE: func(cmd *cobra.Command, args []string) error {
		err:=buildFrontendCommand.RunE(cmd,args)
		if err!=nil {
			return err
		}

		err=buildBackendCommand.RunE(cmd,args)
		if err!=nil {
			return err
		}
		return nil
	},
}