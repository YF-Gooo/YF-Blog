package service

import (
	"github.com/gin-gonic/gin"
	"singo/serializer"
	"os"
	"io/ioutil"
	"fmt"
)
// 列出图片
func ListImageService(c *gin.Context) serializer.Response{
	files, err := ioutil.ReadDir(os.Getenv("IMAGE_STORE"));
	if err != nil {
		return serializer.Response{
			Status: 50005,
			Msg:    "Failed to list data into file",
			Error:  err.Error(),
		}
	}
	fmt.Println(files)
	return serializer.Response{
		Data: serializer.BuildImages(files),
	}
}