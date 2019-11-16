package service

import (
	"singo/model"
	"singo/serializer"
	"fmt"
)

// DeleteArticleService 删除投稿的服务
type DeleteArticleService struct {
}

// Delete 删除文章
func (service *DeleteArticleService) Delete(user string, id string) serializer.Response {
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
	fmt.Println(user != Article.UserName)
	if user != Article.UserName {
		return serializer.Response{
			Status: 404,
			Msg:    "没有权限删除",
			Error:  "no permission",
		}
	}
	err = model.DB.Delete(&Article).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "文章删除失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{}
}
