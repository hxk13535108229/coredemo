package http

import "github.com/gohade/hade/framework/gin"

func NewHttpEngine() (*gin.Engine,error) {
	gin.SetMode(gin.ReleaseMode)
	// r:=gin.Default()
	r:=gin.New()
	r.Use(gin.Recovery())

	Routes(r)
	return r,nil
}