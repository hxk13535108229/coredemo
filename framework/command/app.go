package command

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gohade/hade/framework/cobra"
	"github.com/gohade/hade/framework/contract"
)

func initAppCommand() *cobra.Command {
	appCommand.AddCommand(appStartCommand)
	return appCommand
}

var appCommand = &cobra.Command{
	Use: "app",
	Short: "业务应用控制命令",
	Long: "业务应用控制命令，其包括业务启动，关闭，重启，查询等功能",
	RunE: func(cmd *cobra.Command, args []string) error {
		//打印帮助文档
		cmd.Help()
		return nil
	},
}

var appStartCommand =&cobra.Command{
	Use: "start",
	Short: "启动一个Web服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		contaniner:=cmd.GetContainer()
		kernelService := contaniner.MustMake(contract.KernelKey).(contract.Kernel)
		core:=kernelService.HttpEngine()

		//创建一个Server服务
		server:= &http.Server{
			Handler: core,
			Addr: ":8888",
		}

		go func ()  {
			server.ListenAndServe()
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM,syscall.SIGQUIT)
		<-quit

		timeoutCtx,cancel := context.WithTimeout(context.Background(),5*time.Second)
		defer cancel()

		if err:=server.Shutdown(timeoutCtx);err!=nil{
			log.Fatal("Server Shutdown:",err)
		}
		return nil
	},
}