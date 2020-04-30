package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/etcdv3"
)

func main() {
	etcdRegistry := etcdv3.NewRegistry()
	ginRouter := gin.Default()
	ginRouter.Handle("GET", "/user", func(ctx *gin.Context) {
		ctx.String(200, "user api")
	})

	ginRouter.Handle("GET", "/news", func(ctx *gin.Context) {
		ctx.String(200, "news api")
	})

	server := web.NewService(
		web.Registry(etcdRegistry),
		web.Name("testServer"),
		web.Address(":8000"),
		web.Handler(ginRouter),
	)

	server.Run()
}
