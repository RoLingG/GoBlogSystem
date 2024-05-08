package qiniu

import (
	"GoRoLingG/config"
	"GoRoLingG/global"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"time"
)

// 获取七牛云上传的toke
func getToken(q config.QiNiu) string {
	accessKey := q.AccessKey
	secretKey := q.SecretKey
	bucket := q.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

// 获取七牛云上传的配置
func getConfig(q config.QiNiu) storage.Config {
	cfg := storage.Config{}
	//空间对应的机房
	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone = &zone
	//是否使用https域名
	cfg.UseHTTPS = false
	//上传是否使用CDN加速
	cfg.UseCdnDomains = false
	return cfg
}

// uploadImages 七牛云上传图片
func UploadImages(data []byte, imageName string, prefix string) (filePath string, err error) {
	if !global.Config.QiNiu.Isenable {
		return "", errors.New("请启用七牛云上传")
	}
	q := global.Config.QiNiu
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", errors.New("请配置AccessKey和SecretKey")
	}
	if float64(len(data))/1024/1024 > q.Size {
		return "", errors.New("上传图片超过设定大小，请重传")
	}
	upToken := getToken(q)
	cfg := getConfig(q)

	formUploadr := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))

	//获取当前时间
	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf("%s/%s_%s", prefix, now, imageName)
	err = formUploadr.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", q.CDN, ret.Key), nil
}
