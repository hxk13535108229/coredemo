package config

import (
	"path/filepath"

	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
)

type HadeConfigProvider struct {

}

func (provider *HadeConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewHadeConfig
}

func (provider *HadeConfigProvider) Boot(c framework.Container) error {
	return nil
}

func (provider *HadeConfigProvider) IsDefer() bool {
	return false
}

func (provider *HadeConfigProvider) Params(c framework.Container) []interface{} {
	appService := c.MustMake(contract.AppKey).(contract.App)
	envService := c.MustMake(contract.EnvKey).(contract.Env)
	env:=envService.AppEnv()
	//配置文件夹地址
	configFolder:=appService.ConfigFolder()
	envFolder :=filepath.Join(configFolder,env)
	return []interface{} {c,envFolder,envService.All()}
}

func (provider *HadeConfigProvider) Name()string {
	return contract.ConfigKey
}