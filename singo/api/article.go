package api

import (
	service "singo/service/article"
	"github.com/gin-gonic/gin"
)


// CreateArticle 文章投稿
func CreateArticle(c *gin.Context) {
	service := service.CreateArticleService{}
	if err := c.ShouldBind(&service); err == nil {
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
	service := service.UpdateArticleService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteVideo 删除视频的接口
func DeleteArticle(c *gin.Context) {
	service := service.DeleteArticleService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}
