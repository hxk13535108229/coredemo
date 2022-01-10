package demo

import (
	"fmt"

	"github.com/gohade/hade/framework"
)

//服务提供方
type DemoServiceProvider struct{

}

func (sp *DemoServiceProvider) Name() string{
	return Key
}

func (sp *DemoServiceProvider) Register(c framework.Container) framework.NewInstance{
	return NewDemoService
}

func (sp *DemoServiceProvider) IsDefer() bool{
	return true
}

func (sp *DemoServiceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

func (sp *DemoServiceProvider) Boot(c framework.Container) error{
	fmt.Println("demo service boot")
	return nil
}