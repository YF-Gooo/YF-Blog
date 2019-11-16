package service

import (
	"singo/model"
	"singo/serializer"
	"fmt"
)

// UpdateArticleService 更新文章的服务
type UpdateArticleService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"max=300"`
	Markdown  string `form:"markdown" json:"markdown" binding:"max=100000"`
}

// Update 更新文章
func (service *UpdateArticleService) Update(user string,id string) serializer.Response {
	var Article model.Article
	err := model.DB.First(&Article, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "文章不存在",
			Error:  err.Error(),
		}
	}
	fmt.Println(user)
	fmt.Println(Article.UserName)
	if user!=Article.UserName {
		return serializer.Response{
			Status: 404,
			Msg:    "没有权限更改",
			Error:  "no permission",
		}
	}
	Article.Title = service.Title
	Article.Info = service.Info
	Article.Markdown = service.Markdown
	err = model.DB.Save(&Article).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "文章保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildArticle(Article),
	}
}
