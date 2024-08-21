package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
)

type Config struct {
	Endpoint        string
	Bucket          string
	AccessKeyID     string
	AccessKeySecret string
}

var conf *Config

func LoadOSSConfig(config *Config) {
	conf = &Config{
		Endpoint:        "https://" + config.Endpoint,
		Bucket:          config.Bucket,
		AccessKeyID:     config.AccessKeyID,
		AccessKeySecret: config.AccessKeySecret,
	}
}

func UploadFile(key string, file multipart.File) error {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New(conf.Endpoint, conf.AccessKeyID, conf.AccessKeySecret)
	if err != nil {
		return err
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket(conf.Bucket)
	if err != nil {
		return err
	}

	err = bucket.PutObject(key, file)
	if err != nil {
		return err
	}

	return nil
}
