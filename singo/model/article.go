package model

import "github.com/jinzhu/gorm"

// Article 文章模型
type Article struct {
	gorm.Model
	Title string
	Info  string
	UserName  string
	Tag   string
	Content string
	Public bool
}
