package serializer

import (
	"os"
)

// Image 视频序列化器
type Image struct {
	Name     string `json:"name"`
}
// BuildImage 序列化视频
func BuildImage(item os.FileInfo) Image {
	return Image{
		Name: item.Name(),
	}
}
// BuildImages 序列化视频列表
func BuildImages(items []os.FileInfo) (images []Image) {
	for _, item := range items {
		image := BuildImage(item)
		images = append(images, image)
	}
	return images
}