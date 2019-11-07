package service

import (
	"singo/model"
	"singo/serializer"
)

// DeleteArticleService 删除投稿的服务
type DeleteArticleService struct {
}

// Delete 删除文章
func (service *DeleteArticleService) Delete(id string) serializer.Response {
	var Article model.Article
	err := model.DB.First(&Article, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "文章不存在",
			Error:  err.Error(),
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
