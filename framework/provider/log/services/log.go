package services

import (
	"context"
	"io"
	pkgLog "log"
	"time"

	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/provider/log/formatter"
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
func (log *HadeLog) logf(level contract.LogLevel, ctx context.Context, msg string, fields map[string]interface{}) error {
	//先判断日志级别
	if !log.IsLevelEnable(level) {
		return nil
	}

	//使用ctxFielder获取context中的信息
	fs := fields
	if log.ctxFielder != nil {
		t := log.ctxFielder(ctx)
		if t != nil {
			for k, v := range t {
				fs[k] = v
			}
		}
	}

	//如果绑定了trace服务，获取trace信息
	if log.c.IsBind(contract.TraceKey) {
		tracer:=log.c.MustMake(contract.TraceKey).(contract.Trace)
		tc:=tracer.GetTrace(ctx)
		if tc!=nil {
			maps:= tracer.ToMap(tc)
			for k,v := range maps {
				fs[k]=v
			}
		}
	}

	//将日志信息按照formatter序列化成字符串
	if log.formatter == nil {
		log.formatter = formatter.TextFormatter
	}
	ct, err := log.formatter(level, time.Now(), msg, fs)
	if err != nil {
		return err
	}

	//如果是panic级别，则使用log进行panic
	if level == contract.PanicLevel {
		pkgLog.Panicln(string(ct))
		return nil
	}

	//通过output进行输出
	log.output.Write(ct)
	log.output.Write([]byte("\r\n"))
	return nil
}

func (log *HadeLog) SetOutput(output io.Writer) {
	log.output = output
}

func (log *HadeLog) Panic(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.PanicLevel, ctx, msg, fields)
}

func (log *HadeLog) Fatal(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.FatalLevel, ctx, msg, fields)
}

func (log *HadeLog) Error(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.ErrorLevel, ctx, msg, fields)
}

func (log *HadeLog) Warn(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.WarnLevel, ctx, msg, fields)
}

func (log *HadeLog) Info(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.InfoLevel, ctx, msg, fields)
}

func (log *HadeLog) Debug(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.DebugLevel, ctx, msg, fields)
}

func (log *HadeLog) Trace(ctx context.Context, msg string, fields map[string]interface{}) {
	log.logf(contract.TraceLevel, ctx, msg, fields)
}

func (log *HadeLog) SetLevel(level contract.LogLevel) {
	log.level = level
}

func (log *HadeLog) SetCtxFielder(handler contract.CtxFielder) {
	log.ctxFielder = handler
}

func (log *HadeLog) SetFormatter(formatter contract.Formatter) {
	log.formatter = formatter
}
