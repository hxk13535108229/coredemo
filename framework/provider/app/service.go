package app

import (
	"errors"
	"flag"
	"path/filepath"

	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/util"
)

//代表hade框架的App实现
type HadeApp struct {
	contianer framework.Container //服务容器
	baseFolder string 			// 基础路径
}

func (h HadeApp) Version() string {
	return "0.0.3"
}

func (h HadeApp) BaseFolder() string {
	if h.baseFolder!= "" {
		return h.baseFolder
	}

	//如果没有设置，则使用参数
	var baseFolder string
	flag.StringVar(&baseFolder,"base_folder","","base_folder参数，默认为当前路径")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}

	return util.GetExecDirectory()
}

//表示配置文件地址
func(h HadeApp) ConfigFolder() string {
	return filepath.Join(h.BaseFolder(),"config")
}

//表达日志存放地址
func(h HadeApp) LogFolder() string {
	return filepath.Join(h.BaseFolder(),"log")
}

func(h HadeApp) HttpFolder() string {
	return filepath.Join(h.BaseFolder(),"http")
}

func(h HadeApp) ConsoleFolder() string {
	return filepath.Join(h.BaseFolder(),"console")
}

func(h HadeApp) StorageFolder() string {
	return filepath.Join(h.BaseFolder(),"storage")
}

func(h HadeApp) ProviderFolder() string {
	return filepath.Join(h.BaseFolder(),"provider")
}

func(h HadeApp) MiddlewareFolder() string {
	return filepath.Join(h.HttpFolder(),"middleware")
}

func(h HadeApp) CommandFolder() string {
	return filepath.Join(h.ConsoleFolder(),"command")
}

func(h HadeApp) RuntimeFolder() string {
	return filepath.Join(h.StorageFolder(),"runtime")
}

func(h HadeApp) TestFolder() string {
	return filepath.Join(h.BaseFolder(),"test")
}

//初始化HadeApp
func NewHadeApp(params ...interface{}) (interface{},error) {
	if len(params) !=2 {
		return nil,errors.New("param error")
	}

	contianer:=params[0].(framework.Container)
	baseFolder:=params[1].(string)
	return &HadeApp{
		baseFolder: baseFolder,
		contianer: contianer,
	},nil
}