package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	//1.创建一个默认的路由引擎
	r := gin.Default()
	//2.定义路由
	//2.1GET请求处理传参
	//2.1.1普通参数
	//http://localhost:8080/user/save?id=11&name=zhangsan
	r.GET("/user/save", func(c *gin.Context) {
		id := c.Query("id")
		name := c.Query("name")
		c.JSON(200, gin.H{
			"id":   id,
			"name": name,
		},)
	})
	//如果参数不存在，就给一个默认值：
	r.GET("/user/save", func(ctx *gin.Context) {
			id := ctx.Query("id")
			name := ctx.Query("name")
			address := ctx.DefaultQuery("address", "北京")
			ctx.JSON(200, gin.H{
				"id":      id,
				"name":    name,
				"address": address,
			})
		})
	//判断参数是否存在
	r.GET("/user/save", func(ctx *gin.Context) {
			id, ok := ctx.GetQuery("id")
			address, aok := ctx.GetQuery("address")
			ctx.JSON(200, gin.H{
				"id":      id,
				"idok":    ok,
				"address": address,
				"aok":     aok,
			})
		})
	//id是数值类型，上述获取的都是string类型，根据类型获取
	type User struct {
	Id   int64  `form:"id"`
	Name string `form:"name"`
	}
	r.GET("/user/save", func(ctx *gin.Context) {
			var user User
			err := ctx.BindQuery(&user)
			if err != nil {
				log.Println(err)
			}
			ctx.JSON(200, user)
	})
	//或者
	type User2 struct {
	Id      int64  `form:"id"`
	Name    string `form:"name"`
	Address string `form:"address" binding:"required"`
	}
	r.GET("/user/save", func(ctx *gin.Context) {
			var user User2
			err := ctx.ShouldBindQuery(&user)
			if err != nil {
				log.Println(err)
			}
			ctx.JSON(200, user)
		})
	//2.1.2数组参数
	//http://localhost:8080/user/save?address=Beijing&address=shanghai
	r.GET("/user/save", func(ctx *gin.Context) {
		address := ctx.QueryArray("address")
		ctx.JSON(200, address)
	})
	r.GET("/user/save", func(ctx *gin.Context) {
		address, ok := ctx.GetQueryArray("address")
		fmt.Println(ok)
		ctx.JSON(200, address)
	})
	//2.1.3map参数
	//http://localhost:8080/user/save?addressMap[home]=Beijing&addressMap[company]=shanghai
	r.GET("/user/save", func(ctx *gin.Context) {
		addressMap := ctx.QueryMap("addressMap")
		ctx.JSON(200, addressMap)
	})
	r.GET("/user/save", func(ctx *gin.Context) {
		addressMap, _ := ctx.GetQueryMap("addressMap")
		ctx.JSON(200, addressMap)
	})
	//2.2POST请求处理传参
	type User3 struct {
		Id      int64  `form:"id"`
		Name    string `form:"name"`
		Address string `form:"address" binding:"required"`
		AddressMap map[string]string `form:"addressMap"`
	}
	//表单参数
	r.POST("/user/save", func(ctx *gin.Context) {
			id := ctx.PostForm("id")
			name := ctx.PostForm("name")
			address := ctx.PostFormArray("address")
			addressMap := ctx.PostFormMap("addressMap")
			ctx.JSON(200, gin.H{
				"id":         id,
				"name":       name,
				"address":    address,
				"addressMap": addressMap,
			})
		})
	r.POST("/user/save", func(ctx *gin.Context) {
		var user User3
		err := ctx.ShouldBind(&user)
		addressMap, _ := ctx.GetPostFormMap("addressMap")
		user.AddressMap = addressMap
		fmt.Println(err)
		ctx.JSON(200, user)
	})
	//JSON参数	
	r.POST("/user/save", func(ctx *gin.Context) {
			var user User3
			err := ctx.ShouldBindJSON(&user)
			fmt.Println(err)
			ctx.JSON(200, user)
		})
	//2.3处理请求路径参数
	// http://localhost:8080/user/save/111
	r.POST("/user/save/:id", func(ctx *gin.Context) {
		ctx.JSON(200, ctx.Param("id"))
	})
	//2.4处理请求文件参数
	r.POST("/user/save", func(ctx *gin.Context) {
		form, err := ctx.MultipartForm()
		if err != nil {
			log.Println(err)
		}
		files := form.File
		for _, fileArray := range files {
			for _, v := range fileArray {
				ctx.SaveUploadedFile(v, "./"+v.Filename)
			}

		}
		ctx.JSON(200, form.Value)
	})

}