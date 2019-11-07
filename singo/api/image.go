package api

import (
	service "singo/service/image"
	"github.com/gin-gonic/gin"
)

// UploadImage 上传图片投稿
func UploadImage(c *gin.Context) {
	res :=service.UploadImageService(c)
	c.JSON(200, res)
}

// ListImage 列出图片
func ListImage(c *gin.Context) {
	res :=service.ListImageService(c)
	c.JSON(200, res)
}



