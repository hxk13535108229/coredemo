package cobra

import "github.com/gohade/hade/framework"

//设置服务容器
func (c *Command) SetContainer(container framework.Container) {
	c.contianer=container
}

//获取容器
func (c *Command) GetContainer() framework.Container {
	return c.Root().contianer
}