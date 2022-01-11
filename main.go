package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gohade/hade/app/provider/demo"
	"github.com/gohade/hade/framework/gin"
	"github.com/gohade/hade/framework/middleware"
	"github.com/gohade/hade/framework/provider/app"

	hadeHttp "github.com/gohade/hade/app/http"
)

func main() {
	//创建engine的结构
	core := gin.New()
	core.Bind(&app.HadeAppProvider{})
	core.Bind(&demo.DemoProvider{})

	core.Use(gin.Recovery())
	core.Use(middleware.Cost())

	hadeHttp.Routes(core)

	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}

	go func() {
		server.ListenAndServe()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

}
