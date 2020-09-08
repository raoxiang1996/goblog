package model

type Tag struct {
	ID   int
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}
