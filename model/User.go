package model

import (
	"encoding/base64"
	"goblog/utils/errmsg"
	"log"

	"golang.org/x/crypto/scrypt"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// 查询用户是否存在
func CheckUser(data User) int {
	if data.Username == "" {
		return errmsg.ERROR_UERNAME_EMPTY
	}
	if data.Password == "" {
		return errmsg.ERROR_PASSWORD_EMPTY
	}
	var user User
	db.Select("id").Where("username = ?", data.Username).First(&user)
	if data.ID > 0 {
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
func GetUsers(pageSize int, pageNum int) ([]User, int) {
	var userList []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&userList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR_USER_NOT_EXIST
	}
	return userList, errmsg.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 修改用户
func UpdateUser(id int, data *User) int {
	var user User
	maps := make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&user).Where("id = ?", id).Update(maps)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 密码加密
func (u *User) BeforeSave() {
	u.Password = ScrypyPw(u.Password)
}

func ScrypyPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 11, 222, 11}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}
