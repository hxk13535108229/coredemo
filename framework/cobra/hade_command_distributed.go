package cobra

import (
	"log"
	"time"

	"github.com/gohade/hade/framework/contract"
	"github.com/robfig/cron/v3"
)

func (c *Command) AddDistributedCronCommand(serviceName string,spec string,cmd *Command,holdTime time.Duration) {
	root:=c.Root()

	//初始化cron
	if root.Cron == nil {
		root.Cron = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor)))
		root.CronSpecs=[]CronSpec{}
	}

	//cron命令的指数，这里要注意Type为distributed-cron
	root.CronSpecs=append(root.CronSpecs, CronSpec{
		Type: "distributed-cron",
		Cmd: cmd,
		Spec: spec,
		ServiceName: serviceName,
	})

	appService := root.GetContainer().MustMake(contract.AppKey).(contract.App)
	distributeServce := root.GetContainer().MustMake(contract.DistributedKey).(contract.Distributed)
	appID := appService.AppID()

	//复制要执行的command为cronCmd，并且设置为rootCmd
	var cronCmd Command
	ctx:=root.Context()
	cronCmd = *cmd
	cronCmd.args=[]string{}
	cronCmd.SetParentNull()

	//cron增加匿名函数
	root.Cron.AddFunc(spec,func (){
		//防止panic
		defer func ()  {
			if err:= recover();err!=nil{
				log.Println(err)
			}
		}()

		//节点进行选举，返回选举的结果
		selectedAppID,err:=distributeServce.Select(serviceName,appID,holdTime)
		if err!=nil{
			return 
		}
		//如果自己没有被选择到，直接返回
		if selectedAppID!=appID {
			return 
		}

		//如果被选择到了，执行这个定时的任务
		err = cronCmd.ExecuteContext(ctx)
		if err!=nil{
			log.Println(err)
		}
	})

}

