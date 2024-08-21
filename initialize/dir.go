package initialize

import (
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/pkg/files"
)

func CreateDirectories() {
	var directories []string

	directories = append(directories, global.Config.Log.Dir)
	directories = append(directories, global.Config.Image.AvatarDir)

	for _, dir := range directories {
		files.CreateRootDir(dir)
	}
}
