package app

import (
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
)

//提供App的具体实现方法
type HadeAppProvider struct {
	BaseFolder string
}

//注册HadeApp方法
func (h *HadeAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewHadeApp
}

//启动调用
func (h *HadeAppProvider) Boot(container framework.Container) error{
	return nil
}

func (h *HadeAppProvider) IsDefer() bool {
	return false
}

func (h *HadeAppProvider) Params(container framework.Container) []interface{} {
	return []interface{} {container,h.BaseFolder}
}

func (h *HadeAppProvider) Name() string {
	return contract.AppKey
}