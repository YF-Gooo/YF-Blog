package service

import (
	"singo/model"
	"singo/serializer"
)

// ShowArticleService 投稿详情的服务
type ShowArticleService struct {
}

// Show 创建文章
func (service *ShowArticleService) Show(id string) serializer.Response {
	var Article model.Article
	err := model.DB.First(&Article, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "文章不存在",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildArticle(Article),
	}
}
