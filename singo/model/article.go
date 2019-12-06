package model

import "github.com/jinzhu/gorm"

// Article 文章模型
type Article struct {
	gorm.Model
	Title string
	Info  string
	UserName  string
	like int
	view int 
	Tag   string
	Markdown string `sql:"type:text" valid:"required,length(1|10000)"`
	Public bool
}
