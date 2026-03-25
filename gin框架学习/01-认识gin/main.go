package main

import "github.com/gin-gonic/gin"

func main() {
	//1.创建一个默认的路由引擎
	r := gin.Default()
	//2.定义路由
	r.GET("/ping",func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"pong",
		})
	})
	//监听并在0.0.0.0:8080上启动服务,可以在Run()自定义端口,比如Run(":9999")
	r.Run() 
}