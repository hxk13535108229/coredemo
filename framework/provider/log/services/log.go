package services

import (
	"context"
	"io"

	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
)

//通用的日志实例
type HadeLog struct {
	//五个必要的参数
	//日志级别
	level contract.LogLevel

	//日志的格式化方式
	formatter contract.Formatter

	//ctx获取上下文字段
	ctxFielder contract.CtxFielder

	//输出
	output io.Writer

	//容器
	c framework.Container
}

//判断这个级别是否可以打印
func (log *HadeLog) IsLevelEnable(level contract.LogLevel) bool {
	return level <= log.level
}

//打印日志的核心函数
func (log *HadeLog) logf(level contract.LogLevel,ctx context.Context,msg string,fields map[string]interface{}) error{
	//先判断日志级别
	if !log.IsLevelEnable(level) {
		return nil
	}

	//使用ctxFielder获取context中的信息
	fs:=fields
	if log.ctxFielder!=nil{
		t:=log.ctxFielder(ctx)
		if t !=nil {
			for k,v:=range t {
				fs[k]=v
			}
		}
	}

	//如果绑定了trace服务，获取trace信息
	// if log.c.IsBind(contract.)
}