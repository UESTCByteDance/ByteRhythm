package util

import (
	"ByteRhythm/config"
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func UploadVideo(data []byte) (VideoUrl string, err error) {
	config.Init()
	size := int64(len(data))
	key := fmt.Sprintf("%s.mp4", GenerateUUID())
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", config.Bucket, key),
	}
	mac := qbox.NewMac(config.AccessKey, config.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	uploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err = uploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), size, &putExtra)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", config.Domain, ret.Key), nil
}
