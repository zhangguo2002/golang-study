package main

import "github.com/gin-gonic/gin"

func main() {
	//1.初始话路由引擎
	r := gin.Default()
	//2.定义路由
	r.GET("/cookie", func(c *gin.Context){
		//设置cookie
		c.SetCookie("site_cookie", "cookievalue", 3600, 
		"/", "localhost", false, true)
	})
	r.GET("/read", func(c *gin.Context) {
		// 根据cookie名字读取cookie值
		data, err := c.Cookie("site_cookie")
		if err != nil {
			// 直接返回cookie值
			c.String(200,data)
			return
		}
		c.String(200,"not found!")
	})
	r.GET("/del", func(c *gin.Context) {
		// 设置cookie  MaxAge设置为-1，表示删除cookie
		c.SetCookie("site_cookie", "cookievalue", -1, "/", "localhost", false, true)
		c.String(200,"删除cookie")
	})

}