package http

import (
	"github.com/gohade/hade/app/http/module/demo"
	"github.com/gohade/hade/framework/gin"
	"github.com/gohade/hade/framework/middleware"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/","./dist/")

	r.Use(middleware.Trace())
	demo.Register(r)
}