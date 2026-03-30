package api

import (
	"time"

	"gormtest/dao"

	"github.com/gin-gonic/gin"
)

// RegisterRouter 统一注册所有路由
func RegisterRouter(r *gin.Engine) {
	r.GET("/save", SaveUser)
	r.GET("/user", GetUser)
	r.GET("/users", GetAllUsers)
	r.GET("/update", UpdateUser)
	r.GET("/delete", DeleteUser)
}

func SaveUser(c *gin.Context) {
	user := &dao.User{
		Username:   "zhangsan",
		Password:   "123456",
		CreateTime: time.Now().UnixMilli(),
	}
	dao.Save(user)
	c.JSON(200, user)
}

// 获取单个用户 (这里为测试方便写死ID为1)
func GetUser(c *gin.Context) {
	user := dao.GetById(1)
	c.JSON(200, user)
}

// 获取所有用户
func GetAllUsers(c *gin.Context) {
	users := dao.GetAll()
	c.JSON(200, users)
}

// 更新用户 (这里为测试方便写死ID为1)
func UpdateUser(c *gin.Context) {
	dao.UpdateById(1)
	user := dao.GetById(1)
	c.JSON(200, user)
}

// 删除用户 (这里为测试方便写死ID为1)
func DeleteUser(c *gin.Context) {
	dao.DeleteById(1)
	c.JSON(200, gin.H{"message": "用户删除成功"})
}