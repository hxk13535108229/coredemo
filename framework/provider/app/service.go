package app

import (
	"errors"
	"flag"
	"path/filepath"

	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/util"
	"github.com/google/uuid"
)

//代表hade框架的App实现
type HadeApp struct {
	contianer  framework.Container //服务容器
	baseFolder string              // 基础路径
	appId      string              //当前这个app的唯一id，可以用于分布式锁

	//配置加载
	configMap map[string]string
}

func (app HadeApp) Version() string {
	return "0.0.3"
}

func (app HadeApp) BaseFolder() string {
	if app.baseFolder != "" {
		return app.baseFolder
	}

	//如果参数也没有
	return util.GetExecDirectory()
}

//表示配置文件地址
func (app HadeApp) ConfigFolder() string {
	if val,ok:= app.configMap["config_folder"];ok {
		return val
	}

	return filepath.Join(app.BaseFolder(), "config")
}

//表达日志存放地址
func (app HadeApp) LogFolder() string {
	if val,ok:=app.configMap["log_folder"];ok{
		return val
	}
	return filepath.Join(app.StorageFolder(), "log")
}

func (app HadeApp) HttpFolder() string {
	if val,ok :=app.configMap["http_folder"];ok {
		return val
	}
	return filepath.Join(app.BaseFolder(),"app","http")
}

func (app HadeApp) ConsoleFolder() string {
	if val,ok:=app.configMap["console_folder"]; ok{
		return val
	}
	return filepath.Join(app.BaseFolder(), "app","console")
}

func (app HadeApp) StorageFolder() string {
	if val,ok:=app.configMap["storage_folder"];ok{
		return val
	}
	return filepath.Join(app.BaseFolder(), "storage")
}

func (app HadeApp) ProviderFolder() string {
	if val,ok:=app.configMap["provider_folder"];ok{
		return val
	}
	return filepath.Join(app.BaseFolder(),"app", "provider")
}

func (app HadeApp) MiddlewareFolder() string {
	if val,ok:=app.configMap["middleware_folder"];ok{
		return val
	}
	return filepath.Join(app.HttpFolder(), "middleware")
}

func (app HadeApp) CommandFolder() string {
	if val,ok:= app.configMap["command_folder"];ok{
		return val
	}
	return filepath.Join(app.ConsoleFolder(), "command")
}

func (app HadeApp) RuntimeFolder() string {
	if val,ok:=app.configMap["runtime_folder"];ok {
		return val
	}
	return filepath.Join(app.StorageFolder(), "runtime")
}

func (app HadeApp) TestFolder() string {
	if val,ok :=app.configMap["test_folder"];ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "test")
}

//初始化HadeApp
func NewHadeApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	contianer := params[0].(framework.Container)
	baseFolder := params[1].(string)

	//如果没有设置，则使用参数
	if baseFolder == "" {
		flag.StringVar(&baseFolder,"base_folder","","base_folder参数,默认为当前路径")
		flag.Parse()
	}
	appId:=uuid.New().String()
	configMap:=map[string]string {}

	return &HadeApp{
		baseFolder: baseFolder,
		contianer: contianer,
		appId: appId,
		configMap: configMap,
	},nil
}

func (h HadeApp) AppID() string {
	return h.appId
}

func (app *HadeApp) LoadAppConfig(kv map[string]string) {
	for key,val :=range kv {
		app.configMap[key]=val
	}
}

func (app *HadeApp) AppFolder() string {
	if val,ok := app.configMap["app_folder"];ok {
		return val
	}
	return filepath.Join(app.BaseFolder(),"app")
}

func (app *HadeApp) DeployFolder() string {
	if val,ok := app.configMap["deploy_folder"];ok {
		return val
	}
	return filepath.Join(app.BaseFolder(),"deploy")
}