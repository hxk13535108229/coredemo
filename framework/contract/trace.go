package contract

import (
	"context"
	"net/http"
)

const TraceKey="hade:trace"

const(
	TraceKeyTraceID ="trace_id"
	TraceKeySpanID="span_id"
	TraceKeyCspanID="cspan_id"
	TraceKeyParentID="parent_id"
	TraceKeyMethod="method"
	TraceKeyCaller="caller"
	TraceKeyTime="time"
)

type TraceContext struct {
	//global unique
	TraceID string

	//父节点SpanID
	ParentID string

	//当前节点SpanID
	SpanID string

	//子节点调用的SpanID，由调用方指定
	CspanID string

	//标记各种信息
	Annotation map[string]string
}

type Trace interface {
	WithTrace(c context.Context,trace *TraceContext) context.Context

	GetTrace(c context.Context) *TraceContext

	NewTrace() *TraceContext

	StartSpan(trace *TraceContext) *TraceContext

	ToMap(Trace *TraceContext) map[string]string

	ExtractHTTP(req *http.Request) *TraceContext

	InjectHTTP(req *http.Request,trace *TraceContext) *http.Request
}