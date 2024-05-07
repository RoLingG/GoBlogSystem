package images_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/models/ctype"
	"GoRoLingG/plugins/qiniu"
	"GoRoLingG/res"
	"GoRoLingG/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
)

var (
	//图片上传白名单
	WhiteImagesList = []string{
		".jpg",
		".png",
		".jpeg",
		".gif",
		".webp",
		".svg",
		".ico",
	}
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"code"`
	Msg       string `json:"msg"`
}

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

	var imageMsg string
	var resList []FileUploadResponse
	for _, file := range fileList {
		//判断图片后缀是否为白名单内的后缀
		fileName := file.Filename
		//获取图片后缀
		suffix := strings.ToLower(path.Ext(fileName))
		if !utils.InList(suffix, WhiteImagesList) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "图片后缀非法",
			})
			continue
		}

		//获取图片上传完整路径(包括图片名)
		filePath := path.Join(yamlPath, file.Filename)
		//判断图片大小
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.ImagesUpload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片过大,当前图片大小为: %.2f, 图片上传设定大小为: %dMB", size, global.Config.ImagesUpload.Size),
			})
			continue
		}
		//获取图片对象，并对对象byte化进行md5加密
		fileObj, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		byteData, err := io.ReadAll(fileObj)
		imageHash := utils.Md5(byteData)
		//通过哈希的方式，去数据库中查图片是否存在
		var imageModel models.ImageModel
		err = global.DB.Take(&imageModel, "hash = ?", imageHash).Error
		if err == nil {
			//找到了
			resList = append(resList, FileUploadResponse{
				FileName:  imageModel.Path,
				IsSuccess: true,
				Msg:       "图片重复传输,图片入库传输无效",
			})
			imageMsg = "图片重复传输,图片入库传输无效"
			continue
		}

		//判断是否将图片存储于七牛云
		if global.Config.QiNiu.Isenable {
			filePath, err = qiniu.UploadImages(byteData, fileName, "goblogImage")
			if err != nil {
				global.Log.Error(err)
				continue
			}
			resList = append(resList, FileUploadResponse{
				FileName:  filePath,
				IsSuccess: true,
				Msg:       "图片于七牛云上传成功",
			})
			//图片入库
			global.DB.Create(&models.ImageModel{
				Model:     models.Model{},
				Path:      filePath,
				Hash:      imageHash,
				Name:      fileName,
				ImageType: ctype.QiNiu,
			})
			continue
		}

		//图片大小没问题，则判断图片存储是否成功
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			//图片大小有问题
			global.Log.Error(err)
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       err.Error(),
			})
			continue
		}

		//图片上传没问题，则执行下面操作
		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "图片上传成功",
		})
		//图片入库
		global.DB.Create(&models.ImageModel{
			Model:     models.Model{},
			Path:      filePath,
			Hash:      imageHash,
			Name:      fileName,
			ImageType: ctype.Local,
		})
		imageMsg = "图片上传成功"
	}

	res.OKWithDataAndMsg(resList, imageMsg, c)
}
