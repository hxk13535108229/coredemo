package contract

import (
	"context"
	"io"
	"time"
)

const LogKey= "hade:log"

type LogLevel uint32

const(
	//表示位置的日志级别
	UnknownLevel LogLevel =iota

	//表示会导致整个程序出现崩溃的日志信息
	PanicLevel

	//表示会导致当前这个请求出现提前终止的错误信息
	FatalLevel

	//表示出现错误，但是不一定影响后续请求逻辑的报警信息
	ErrorLevel

	//表示出现错误，但是一定不影响后续请求逻辑的报警信息
	WarnLevel

	//表示正常的日志信息输出
	InfoLevel

	//表示在调试状态下打印出来的日志信息
	DebugLevel

	//表示最详细的信息
	TraceLevel
)

//定义了从context中获取信息的方法
type CtxFielder func(ctx context.Context) map[string]interface{}

//定义了将日志信息组织成字符串的通用方法
type Formatter func(level LogLevel,t time.Time,msg string,fields map[string]interface{}) ([]byte,error)

//日志服务协议
type Log interface{
	Panic(ctx context.Context,msg string,fields map[string]interface{})

	Fatal(ctx context.Context,msg string,fields map[string]interface{})

	Error(ctx context.Context,msg string,fields map[string]interface{})

	Warn(ctx context.Context,msg string,fields map[string]interface{})

	Info(ctx context.Context,msg string,fields map[string]interface{})

	Debug(ctx context.Context,msg string,fields map[string]interface{})

	Trace(ctx context.Context,msg string,fields map[string]interface{})

	//设置日志级别
	SetLevel(level LogLevel)

	//从context中获取上下文字段field
	SetCtxFielder(handler CtxFielder)

	//设置输出格式
	SetFormatter(formatter Formatter)

	//设置输出管道
	SetOutput(out io.Writer)
}