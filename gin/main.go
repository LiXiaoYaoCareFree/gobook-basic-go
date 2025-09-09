package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	server.Use(func(ctx *gin.Context) {
		println("这是第一个 Middleware")
	}, func(ctx *gin.Context) {
		println("这是第二个 Middleware")
	})
	//路由注册
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, world")
	})

	//参数路由，路径参数
	server.GET("/users/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "hello, "+name)
	})

	//查询参数
	// GET /order?id=123
	server.GET("/order", func(ctx *gin.Context) {
		id := ctx.Query("id")
		ctx.String(http.StatusOK, "订单 ID 是 "+id)
	})

	server.GET("/views/*.html", func(ctx *gin.Context) {
		view := ctx.Param(".html")
		ctx.String(http.StatusOK, "view 是 "+view)
	})

	server.POST("/login", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, login")
	})

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
