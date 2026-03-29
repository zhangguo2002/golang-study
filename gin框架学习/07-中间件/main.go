package main

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	// 通过use设置全局中间件
	// 设置日志中间件，主要用于打印请求日志
	r.Use(gin.Logger())
	// 设置Recovery中间件，主要用于拦截paic错误，不至于导致进程崩掉
	r.Use(gin.Recovery())
	r.GET("/test", func(c *gin.Context) {
		panic(errors.New("test error"))
	})
	r.GET("/hello",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"hello",
		})
	})
	r.Run(":8080")
}