package contract

import (
	"net/http"

)

const KernelKey= "hade:kernel"

//接口提供框架最核心的架构
type Kernel interface {
	HttpEngine() http.Handler
}

