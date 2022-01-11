package kernel

import (
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/gin"
)

//提供web引擎
type HadeKernelProvider struct {
	HttpEngine *gin.Engine
}

//注册服务提供者
func (provider *HadeKernelProvider) Register(c framework.Container) framework.NewInstance {
	return NewHadeKernelService
}

func (provider *HadeKernelProvider) Boot(c framework.Container) error {
	if provider.HttpEngine == nil {
		provider.HttpEngine=gin.Default()
	}
	provider.HttpEngine.SetContainer(c)
	return nil
} 

func (provider *HadeKernelProvider) IsDefer() bool {
	return false
}

func (provider *HadeKernelProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.HttpEngine}
}

func (provider *HadeKernelProvider) Name() string {
	return contract.KernelKey
}