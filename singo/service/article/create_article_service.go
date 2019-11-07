package service

import (
	"singo/model"
	"singo/serializer"
)

// CreateArticleService 文章投稿的服务
type CreateArticleService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"max=300"`
	Content  string `form:"content" json:"content" binding:"max=10000"`

}

// Create 创建视频
func (service *CreateArticleService) Create() serializer.Response {
	Article := model.Article{
		Title: service.Title,
		Info:  service.Info,
		Content:  service.Content,
	}

	err := model.DB.Create(&Article).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "视频保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildArticle(Article),
	}
}
