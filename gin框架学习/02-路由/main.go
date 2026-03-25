package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//1.创建一个默认的路由引擎
	r := gin.Default()
	//2.定义路由
	//2.1请求方法
	//get请求 读取服务器上的资源
	r.GET("/get",func (c *gin.Context)  {
		c.JSON(200,gin.H{
		    "message":"get",
		})
	})
	//post请求 在服务器上创建资源
	r.POST("/post",func (c *gin.Context)  {
		c.JSON(200,gin.H{
			"message":"post",
		})
	})
	//put请求 在服务器上更新资源
	r.PUT("/put",func (c *gin.Context)  {
		c.JSON(200,gin.H{
			"message":"put",
		})
	})
	//delete请求 删除服务器上的资源
	r.DELETE("/delete",func (c *gin.Context)  {
		c.JSON(200,gin.H{
			"message":"delete",
		})
	})
	//patch请求 在服务器上更新资源（部分更新）
	r.PATCH("/patch",func (c *gin.Context)  {
		c.JSON(200,gin.H{
			"message":"patch",
		})
	})
	//2.2URL
	//静态url
	r.POST("bolg/cxy",func (c *gin.Context)  {
		c.JSON(200,gin.H{
			"message":"bolg/cxy",
		})
	})
	//路径参数
	r.POST("user/:id",func (c *gin.Context)  {
		param:=c.Param("id")
		c.JSON(200,gin.H{
			"message":param,
		})
	})
	//模糊匹配
	r.POST("user/*id",func (c *gin.Context)  {
		param:=c.Param("id")
		c.JSON(200,gin.H{
			"message":param,
		})
	})
	//3.处理函数
	type HandlerFunc func(*gin.Context)
	//4.分组路由
	usergroup:=r.Group("/user")
	{
		usergroup.GET("/login",func (c *gin.Context)  {
			c.JSON(200,gin.H{
				"message":"login",
			})
		})
		usergroup.GET("/logout",func (c *gin.Context)  {
			c.JSON(200,gin.H{
				"message":"logout",
			})
		})
	}
	r.Run(":8888")

}