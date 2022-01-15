package services

import (
	"os"
	"path/filepath"

	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/util"
	"github.com/pkg/errors"
)

//表示单个日志文件输出
type HadeSingleLog struct{
	HadeLog

	folder string
	file string
	fd *os.File
}

func NewHadeSingleLog(params ...interface{}) (interface{},error) {
	c:=params[0].(framework.Container)
	level:=params[1].(contract.LogLevel)
	ctxFielder:=params[2].(contract.CtxFielder)
	formatter:=params[3].(contract.Formatter)

	appService:=c.MustMake(contract.AppKey).(contract.App)
	configService:=c.MustMake(contract.ConfigKey).(contract.Config)

	log:=&HadeSingleLog{}
	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)
	//从配置文件中获取folder信息，否则使用默认的LogFolder文件夹
	folder:=appService.LogFolder()
	if configService.IsExist("log.folder") {
		folder=configService.GetString("log.folder")
	}
	log.folder=folder
	//如果folder不存在，则创建
	if !util.Exists(folder) {
		os.MkdirAll(folder,os.ModePerm)
	}

	//从配置文件中获取file信息，否则使用默认的hade.log
	log.file="hade.log"
	if configService.IsExist("log.file"){
		log.file=configService.GetString("log.file")
	}


	fd,err:=os.OpenFile(filepath.Join(log.folder,log.file),os.O_APPEND|os.O_CREATE|os.O_RDWR,0666)
	if err!=nil{
		return nil,errors.Wrap(err,"open log file err")
	}
	log.SetOutput(fd)
	log.c=c

	return log,nil
}