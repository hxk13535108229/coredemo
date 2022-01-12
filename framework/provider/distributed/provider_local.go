package distributed

import (
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
)

//提供App的具体实现方法
type LocalDistributedProvider struct {

}

func (h *LocalDistributedProvider) Register(container framework.Container) framework.NewInstance {
	return NewLocalDistributedService
}

func (h *LocalDistributedProvider) Boot(container framework.Container) error{
	return nil
}

func (h *LocalDistributedProvider) IsDefer() bool{
	return false
}


func (h *LocalDistributedProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

func (h *LocalDistributedProvider) Name() string{
	return contract.DistributedKey
}