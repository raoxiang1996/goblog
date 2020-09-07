package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title    string `gorm:"type:varchar(100);not null" json:"title"`
	Execrpt  string `gorm:"type:varchar(200);not null" json:"execrpt"`
	Category Category
	Cid      int    `gorm:"type:int" json:"cid"`
	content  string `gorm:"type:long text" json:"content"`
	Img      string `gorm:"type:varchar(100);not null" json:"img"`
	Tag      Tag
	Tid      int `gorm:"type:int" json:"tid"`
}
