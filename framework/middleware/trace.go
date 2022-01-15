package middleware

import (
	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/gin"
)

//将协程中的函数异常进行捕获
func Trace() gin.HandlerFunc {
	return func (c *gin.Context)  {
		//start time
		tracer := c.MustMake(contract.TraceKey).(contract.Trace)
		traceCtx := tracer.ExtractHTTP(c.Request)

		tracer.WithTrace(c,traceCtx)

		//执行具体的业务逻辑
		c.Next()
	}
}