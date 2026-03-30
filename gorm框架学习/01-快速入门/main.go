package main

import (
	"gormtest/api"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 初始化 Gin 引擎
	r := gin.Default()

	// 2. 将引擎实例传给 API 层的路由注册函数
	api.RegisterRouter(r)

	// 3. 启动 HTTP 服务，默认在 8080 端口监听
	// 注意：只要导入了 gormtest/api 也就是间接导入了 gormtest/dao
	// dao/db.go 中的 init() 函数会自动执行，完成数据库连接。
	r.Run(":8080")
}