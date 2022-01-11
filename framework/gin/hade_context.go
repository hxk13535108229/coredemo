// Copyright 2021 Hade.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package gin

import (
	"context"

)

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
}

//context实现container的几个封装
func (ctx *Context) Make(key string) (interface{}, error){
	return ctx.container.Make(key)
}

func (ctx *Context) MustMake(key string) (interface{}) {
	return ctx.container.MustMake(key)
}

func (ctx *Context) MakeNew(key string, params []interface{}) (interface{},error){
	return ctx.container.MakeNew(key,params)
}
