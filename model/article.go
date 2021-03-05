package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model

	Category Category
	Title    string `gorm:"type:varchar(100);not null" json:"title"`
	Content  string `gorm:"type:longtext;" json:"content"`
	Desc     string `gorm:"type:varchar(200);not null" json:"desc"`
	Cid      int    `gorm:"type:int;not null" json:"cid"` // categoryId
}
