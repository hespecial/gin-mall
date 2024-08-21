package initialize

import (
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/pkg/oss"
)

func InitOSS() {
	if global.Config.Server.UploadMode == "oss" {
		oss.LoadOSSConfig(&oss.Config{
			Endpoint:        global.Config.Oss.Endpoint,
			Bucket:          global.Config.Oss.Bucket,
			AccessKeyID:     global.Config.Oss.AccessKeyID,
			AccessKeySecret: global.Config.Oss.AccessKeySecret,
		})
	}
}
