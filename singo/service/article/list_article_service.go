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
	total:=0
	if  service.Limit==0{
		service.Limit = 6
	}
	if	err := model.DB.Model(model.Article{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}
	if	err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&Articles).Error;err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接错误",
				Error:  err.Error(),
			}
		}
	return serializer.BuildListResponse(serializer.BuildArticles(Articles),uint(total))
}


// UserList 用户文章列表
func (service *ListArticleService) UserList(user string) serializer.Response {
	var Articles []model.Article
	total:=0
	if  service.Limit==0{
		service.Limit = 6
	}
	if	err := model.DB.Where("user_name = ?", user).Model(model.Article{}).Count(&total).Error; err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接错误",
				Error:  err.Error(),
			}
		}
	if err := model.DB.Where("user_name = ?", user).Limit(service.Limit).Offset(service.Start).Find(&Articles).Error ; err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接错误",
				Error:  err.Error(),
			}
		}
	return serializer.BuildListResponse(serializer.BuildArticles(Articles),uint(total))
}

// SearchList 搜索文章列表
func (service *ListArticleService) SearchList(kw string) serializer.Response {
	var Articles []model.Article
	total:=0
	if  service.Limit==0{
		service.Limit = 6
	}
	if	err := model.DB.Where("title LIKE ?","%"+kw+"%").Model(model.Article{}).Count(&total).Error; err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接错误",
				Error:  err.Error(),
			}
		}
	if err := model.DB.Where("title LIKE ?","%"+kw+"%").Limit(service.Limit).Offset(service.Start).Find(&Articles).Error ; err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接错误",
				Error:  err.Error(),
			}
		}
	return serializer.BuildListResponse(serializer.BuildArticles(Articles),uint(total))
}

// UserSearchList 搜索文章列表
func (service *ListArticleService) UserSearchList(user string,kw string) serializer.Response {
	var Articles []model.Article
	total:=0
	if  service.Limit==0{
		service.Limit = 6
	}
	if	err := model.DB.Where("user_name = ?", user).Where("title LIKE ?","%"+kw+"%").Model(model.Article{}).Count(&total).Error; err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接错误",
				Error:  err.Error(),
			}
		}
	if err := model.DB.Where("user_name = ?", user).Where("title LIKE ?","%"+kw+"%").Limit(service.Limit).Offset(service.Start).Find(&Articles).Error ; err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接错误",
				Error:  err.Error(),
			}
		}
	return serializer.BuildListResponse(serializer.BuildArticles(Articles),uint(total))
}
