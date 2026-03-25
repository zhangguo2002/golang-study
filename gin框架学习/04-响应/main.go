package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//1.创建一个默认的路由引擎
	r := gin.Default()
	//2.定义路由
	//字符串响应
	r.GET("/user/save", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "this is a %s", "ms string response")
	})
	//JSON响应
	r.GET("/user/save", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})
	//XML响应
	type XmlUser struct {
	Id   int64  `xml:"id"`
	Name string `xml:"name"`
	}
	r.GET("/user/save", func(ctx *gin.Context) {
			u := XmlUser{
				Id:   11,
				Name: "zhangsan",
			}
			ctx.XML(http.StatusOK, u)
		})
	//文件格式响应
	r.GET("/user/save", func(ctx *gin.Context) {
		//ctx.File("./1.png")
		ctx.FileAttachment("./1.png", "2.png")
	})
	// 设置http响应头
	r.GET("/user/save", func(ctx *gin.Context) {
		ctx.Header("test", "headertest")
	})
	// 重定向响应
	r.GET("/user/save", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	//YAML响应
	r.GET("/user/save", func(ctx *gin.Context) {
		ctx.YAML(200, gin.H{"name": "ms", "age": 19})
	})

}