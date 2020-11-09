package qiniu

import (
	"context"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"io"
	"time"
)

const (
	bucket    = ""
	accessKey = ""
	secretKey = ""
)

func init() {

}

func config() storage.Config {
	cfg := storage.Config{}

	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	return cfg
}

func Upload(localFile io.Reader, size int64, filename string) (string, error) {

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cig := config()
	formUploader := storage.NewFormUploader(&cig)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	err := formUploader.Put(context.Background(), &ret, upToken, filename, localFile, size, &putExtra)
	if err != nil {
		return "", err
	}
	return ret.Key, nil
}

//下载凭证
func QiNiuDownload(fileName string, deadline int64) string {
	if deadline == 0 {
		deadline = time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	}
	mac := qbox.NewMac(accessKey, secretKey)
	domain := bucket
	key := fileName
	privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)
	return privateAccessURL
}
