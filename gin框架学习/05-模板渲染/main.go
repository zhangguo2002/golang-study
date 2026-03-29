package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()
	// // 模板解析
	// r.LoadHTMLFiles("templates/index.tmpl")

	// r.GET("/index", func(c *gin.Context) {
	// 	// HTML请求
	// 	// 模板的渲染
	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"title": "hello 模板",
	// 	})
	// })

	// r.Run(":9090") // 启动server
	//创建一个默认的路由引擎
	r:=gin.Default()
	
	//模板解析
	r.LoadHTMLFiles("templates/index.html")

	//定义路由
	r.GET("/index",func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{
			"title":"hello 模板",
		})
	})
	r.Run(":7979")

}