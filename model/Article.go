package model

import (
	"github.com/jinzhu/gorm"
	"goblog/utils/errmsg"
)

type Article struct {
	gorm.Model
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Execrpt  string   `gorm:"type:varchar(200);not null" json:"execrpt"`
	Category Category `gorm:"foreignkey:cid"`
	Cid      int      `gorm:"type:int" json:"cid"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Img      string   `gorm:"type:varchar(100);not null" json:"img"`
}

// 添加文章
func CreateArticle(data *Article) int {
	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, errmsg.ERROR_GET_ARTICLE_FAIL
	}
	return article, errmsg.SUCCESS
}

// 查询分类下的文章
func GetCateArticle(id int, pageSize int, pageNum int) ([]Article, int) {
	var cateArticleList []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&cateArticleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return cateArticleList, errmsg.SUCCESS
}

// 查询文章列表
func GetArticle(pageSize int, pageNum int) ([]Article, int) {
	var articleList []Article
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return articleList, errmsg.SUCCESS
}

// 删除文章
func DeleteArticle(id int) int {
	var article Article
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 修改文章
func UpdateArticle(id int, data *Article) int {
	var article Article
	maps := make(map[string]interface{})
	maps["title"] = data.Title
	maps["execrpt"] = data.Execrpt
	maps["content"] = data.Content
	maps["img"] = data.Img
	maps["cid"] = data.Cid
	err := db.Model(&article).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// todo 查询分类所有文章

// todo 查询单个文章
