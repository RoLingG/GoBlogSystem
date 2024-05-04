package images_api

import (
	"GoRoLingG/global"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
	"path"
)

// 上传图片，返回图片url
func (ImagesApi) ImagesUploadView(c *gin.Context) {
	Imagesform, err := c.MultipartForm()
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	//Post参数名设置
	fileList, ok := Imagesform.File["images"]
	if !ok {
		res.FailWithMsg("图片不存在", c)
		return
	}
	for _, file := range fileList {
		filePath := path.Join("upload", file.Filename)
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err)
			return
		}
	}
	res.OKWithoutData(c)
}
