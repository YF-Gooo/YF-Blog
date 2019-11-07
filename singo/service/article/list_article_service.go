package service

import (
	"singo/model"
	"singo/serializer"
)

// ListArticleService 文章列表服务
type ListArticleService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 文章列表
func (service *ListArticleService) List() serializer.Response {
	var Articles []model.Article

	err := model.DB.Find(&Articles).
		Limit(service.Limit).Offset(service.Start).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildArticles(Articles),
	}
}
