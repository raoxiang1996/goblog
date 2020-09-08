package model

import (
	"github.com/jinzhu/gorm"

	"goblog/utils/errmsg"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 查询分类是否存在
func CheckCategory(data Category) int {
	if data.Name == "" {
		return errmsg.ERROR_CATEGORY_EMPTY
	}

	var category Category
	db.Select("id").Where("name = ?", data.Name).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED
	}
	return errmsg.SUCCESS
}

// 添加分类
func CreateCategory(data *Category) int {
	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类列表
func GetCategory(pageSize int, pageNum int) ([]Category, int) {
	var categoryList []Category
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categoryList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR_CATEGORY_FAIL
	}
	return categoryList, errmsg.SUCCESS
}

// 删除分类
func DeleteCategory(id int) int {
	var category Category
	err := db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 修改分类
func UpdateCategory(id int, data *Category) int {
	var category Category
	maps := make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&category).Where("id = ?", id).Update(maps)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// todo 查询单个分类下的文章
