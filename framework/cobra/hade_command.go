package cobra

import (
	"log"

	"github.com/gohade/hade/framework"
	"github.com/robfig/cron/v3"
)

//设置服务容器
func (c *Command) SetContainer(container framework.Container) {
	c.contianer=container
}

//获取容器
func (c *Command) GetContainer() framework.Container {
	return c.Root().contianer
}

//保存Cron命令信息，用于展示
type CronSpec struct {
	Type string
	Cmd *Command
	Spec string
	ServiceName string
}

func (c *Command) SetParentNull() {
	c.parent = nil
}

func (c *Command) AddCronCommand(spec string,cmd *Command) {
	//cron结构是挂载在根Command上的
	root:=c.Root()
	if root.Cron == nil {
		//初始化cron
		root.Cron=cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor)))
		root.CronSpecs=[]CronSpec{}
	}
	//增加说明信息
	root.CronSpecs = append(root.CronSpecs,CronSpec{
		Type: "normal-cron",
		Cmd: cmd,
		Spec: spec,
	})

	//制作一个rootCommand
	var cronCmd Command
	ctx:=root.Context()
	cronCmd = *cmd
	cronCmd.args=[]string{}
	cronCmd.SetParentNull()
	cronCmd.SetContainer(root.GetContainer())

	//增加调用函数
	root.Cron.AddFunc(spec,func() {
		//如果后续的command出现panic，这里要捕获
		defer func ()  {
		if err := recover() ;err!=nil{
			log.Println(err)
		}	
		}()

		err := cronCmd.ExecuteContext(ctx)
		if err!=nil {
			log.Println(err)
		}
	})
}