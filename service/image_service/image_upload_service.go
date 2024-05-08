package image_service

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/models/ctype"
	"GoRoLingG/plugins/qiniu"
	"GoRoLingG/utils"
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

var WhiteImagesList = []string{
	".jpg",
	".png",
	".jpeg",
	".gif",
	".webp",
	".svg",
	".ico",
}

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"code"`
	Msg       string `json:"msg"`
}

// 文件上传(主要是上传到数据库)
func (ImageService) ImageUploadService(file *multipart.FileHeader) (uploadRes FileUploadResponse) {
	//判断图片后缀是否为白名单内的后缀
	//先获取图片上传完整路径(包括图片名)
	yamlPath := global.Config.ImagesUpload.Path
	filePath := path.Join(yamlPath, file.Filename)
	fileName := file.Filename
	//↓这里就和之前不一样了，因为本地存储拿的是这个uploadRes.FileName，但之前这个uploadRes.FileName是先定义为filename，再定义为filePath
	//↓但filename的时候已经传过去了，会导致本地物理存储的位置没能定义到filePath指定的路径
	uploadRes.FileName = filePath

	//获取图片后缀
	suffix := strings.ToLower(path.Ext(fileName))
	//对应白名单判断图片后缀是否合法
	if !utils.InList(suffix, WhiteImagesList) {
		uploadRes.Msg = "图片后缀非法"
		return
	}

	//判断图片大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.ImagesUpload.Size) {
		uploadRes.Msg = fmt.Sprintf("图片过大,当前图片大小为: %.2f, 图片上传设定大小为: %dMB", size, global.Config.ImagesUpload.Size)
		return
	}

	//读取文件内容
	//获取图片对象，并对对象byte化进行md5加密
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	//将文件对象byte化
	byteData, err := io.ReadAll(fileObj)
	//将byte化的对象进行MD5加密
	imageHash := utils.Md5(byteData)
	//通过哈希的方式，去数据库中查图片是否存在
	var imageModel models.ImageModel
	err = global.DB.Take(&imageModel, "hash = ?", imageHash).Error
	if err == nil {
		//找到了
		uploadRes.FileName = imageModel.Path
		uploadRes.Msg = "图片已存在,图片入库传输无效"
		return
	}
	//默认上传到本地进行存储
	fileType := ctype.Local
	uploadRes.Msg = "图片于本地上传成功"
	uploadRes.IsSuccess = true

	//判断是否将图片存储于七牛云
	if global.Config.QiNiu.Isenable {
		//qiniu.UploadImages()方法会返回一个七牛云的CDN地址，和一个对应的key
		filePath, err = qiniu.UploadImages(byteData, fileName, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Log.Error(err)
			uploadRes.Msg = err.Error()
			return
		}
		//成功上传，则获取七牛云对应图片的外链
		uploadRes.FileName = filePath
		uploadRes.IsSuccess = true
		uploadRes.Msg = "图片于七牛云上传成功"
		//更改入库类型
		fileType = ctype.QiNiu
	}
	//图片入库
	global.DB.Create(&models.ImageModel{
		Path:      filePath,
		Hash:      imageHash,
		Name:      fileName,
		ImageType: fileType,
	})
	return
}
