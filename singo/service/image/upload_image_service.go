package service

import (
	"github.com/gin-gonic/gin"
	"singo/serializer"
	"bytes"
	"io"
	"singo/util"
	"os"
	"path"
)


// Upload 上传图片
func UploadImageService(c *gin.Context) serializer.Response{
	print(c.Request.FormFile)
	file, head, err := c.Request.FormFile("image")
	FileType:=path.Ext(head.Filename)
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "图片保存失败",
			Error:  err.Error(),
		}
	}
	defer file.Close()

	// 2. 把文件内容转为[]byte
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "Failed to get file data, err:%s\n",
			Error:  err.Error(),
		}
	}
	FileSha1:=util.Sha1(buf.Bytes()) //　计算文件sha1
	// 3. 将文件写入临时存储位置
	Location := os.Getenv("IMAGE_STORE")+FileSha1+FileType // 临时存储地址
	newFile, err := os.Create(Location)
	if err != nil {
		return serializer.Response{
			Status: 50003,
			Msg:    "Failed to create file, err:%s\n",
			Error:  err.Error(),
		}
	}
	defer newFile.Close()
	FileSize:=int64(len(buf.Bytes()))
	nByte, err := newFile.Write(buf.Bytes())
	if int64(nByte) != FileSize || err != nil {
		return serializer.Response{
			Status: 50004,
			Msg:    "Failed to save data into file, writtenSize:%d, err:%s\n",
			Error:  err.Error(),
		}
	}
	
	return serializer.Response{
		Status: 200,
		Msg:    os.Getenv("IMAGE_URL")+FileSha1+FileType,
	}
}



