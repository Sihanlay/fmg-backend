package qiniuUtils

import (
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"grpc-demo/utils"
)

// 获取七牛上传凭证
func GetUploadToken() string {
	putPolicy := storage.PutPolicy{
		Scope:   utils.GlobalConfig.QiNiu.Bucket,
		Expires: utils.GlobalConfig.QiNiu.Expires,
	}
	mac := qbox.NewMac(utils.GlobalConfig.QiNiu.AccessKey, utils.GlobalConfig.QiNiu.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}
