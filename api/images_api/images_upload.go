package images_api

import (
	"GoRoLingG/global"
	"GoRoLingG/res"
	"GoRoLingG/service"
	"GoRoLingG/service/image_service"
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
)

// 上传图片，返回图片url
func (ImagesApi) ImagesUploadView(c *gin.Context) {
	ImagesForm, err := c.MultipartForm()
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	//Post参数名设置
	fileList, ok := ImagesForm.File["images"]
	if !ok {
		res.FailWithMsg("图片不存在", c)
		return
	}

	//判断路径是否存在
	yamlPath := global.Config.ImagesUpload.Path
	_, err = os.ReadDir(yamlPath)
	if err != nil {
		//路径不存在则创建路径
		err = os.MkdirAll(yamlPath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}

	//创建一个list，对每个图片上传进行回复其上传结果
	var resList []image_service.FileUploadResponse

	//for循环轮询post过来的每个图片进行上传判断
	for _, file := range fileList {
		//上传文件，并进行判断，如果成功上传且是七牛云上传，则调用对应方法直接上传到七牛云
		uploadRes := service.Service.ImageService.ImageUploadService(file)
		//如果图片上传失败
		if !uploadRes.IsSuccess {
			resList = append(resList, uploadRes)
			continue
		}

		//如果成功，且是上传到本地，则还要而外进行判断
		if !global.Config.QiNiu.Isenable {
			//ImageUploadService方法将图片上传到数据库，但本地还得保存一份图片
			err = c.SaveUploadedFile(file, uploadRes.FileName)
			if err != nil {
				//图片大小有问题
				global.Log.Error(err)
				resList = append(resList, uploadRes)
				continue
			}
		}
		//如果是成功且是七牛云存储，还得保存
		//(七牛云上传是在service.Service.ImageService.ImageUploadService()这里就已经上传到七牛云了，因此只用存储到resList用于返回给前端对应信息就行)
		resList = append(resList, uploadRes)
	}

	res.OKWithDataAndMsg(resList, "操作成功", c)
}
