package api

import (
	service "singo/service/article"
	"github.com/gin-gonic/gin"
	"singo/model"
	"fmt"
)


// CreateArticle 文章投稿
func CreateArticle(c *gin.Context) {
	service := service.CreateArticleService{}
	if err := c.ShouldBind(&service); err == nil {
		user , _ := c.Get("user")
		if u, ok := user.(*model.User); ok {
			fmt.Println(u.UserName)
			service.UserName = u.UserName
		}
		fmt.Println(service.UserName)
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowArticle 文章详情接口
func ShowArticle(c *gin.Context) {
	service := service.ShowArticleService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// ListArticle 文章列表接口
func ListArticle(c *gin.Context) {
	service := service.ListArticleService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// Update 更新文章的接口
func UpdateArticle(c *gin.Context) {
	user , _ := c.Get("user")
	u, _ := user.(*model.User)
	service := service.UpdateArticleService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(u.UserName,c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteVideo 删除视频的接口
func DeleteArticle(c *gin.Context) {
	user , _ := c.Get("user")
	u, _ := user.(*model.User)
	service := service.DeleteArticleService{}
	res := service.Delete(u.UserName,c.Param("id"))
	c.JSON(200, res)
}
