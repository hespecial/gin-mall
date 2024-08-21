package files

import (
	"os"
)

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func CreateRootDir(dir string) {
	if ok := pathExists(dir); !ok {
		_ = os.MkdirAll(dir, os.ModePerm)
	}
}
