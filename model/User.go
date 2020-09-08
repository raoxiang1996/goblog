package model

import (
	"goblog/utils/errmsg"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// 查询用户是否存在
func CheckUser(user User) int {
	if user.Username == "" {
		return errmsg.ERROR_UERNAME_EMPTY
	}
	if user.Password == "" {
		return errmsg.ERROR_PASSWORD_EMPTY
	}
	db.Select("id").Where("username = ?", user.Username).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_UERNAME_USED
	}
	return errmsg.SUCCESS
}

// 添加用户
func CreateUser(data *User) int {
	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err := db.Limit(pageNum).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 修改用户
func UpdateUser() {

}
