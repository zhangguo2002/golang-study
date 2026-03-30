package dao

import (
	"log"
)

type User struct {
	ID         int64
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	CreateTime int64  `gorm:"column:createtime"`
}

func (u User) TableName() string {
	//绑定MYSQL表名为users
	return "users"
}

func Save(user *User) {
	// 修复：添加 .Error
	err := DB.Create(user).Error
	if err != nil {
		log.Println("insert fail : ", err)
	}
}

func GetById(id int64) User {
	var user User
	err := DB.Where("id=?", id).First(&user).Error
	if err != nil {
		log.Println("get user by id fail : ", err)
	}
	return user
}

func GetAll() []User {
	var users []User
	// 修复：添加 .Error
	err := DB.Find(&users).Error
	if err != nil {
		log.Println("get users  fail : ", err)
	}
	return users
}

func UpdateById(id int64) {
	// 修复：添加 .Error
	err := DB.Model(&User{}).Where("id=?", id).Update("username", "lisi").Error
	if err != nil {
		log.Println("update users  fail : ", err)
	}
}

func DeleteById(id int64) {
	// 修复：添加 .Error
	err := DB.Where("id=?", id).Delete(&User{}).Error
	if err != nil {
		log.Println("delete users  fail : ", err)
	}
}