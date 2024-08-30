package files

import (
	"mime/multipart"
	"net/http"
	"strings"
)

// IsAllowedFileType 检查文件的 MIME 类型是否被允许
func IsAllowedFileType(fileHeader *multipart.FileHeader, allowedTypes []string) bool {
	file, err := fileHeader.Open()
	if err != nil {
		return false
	}
	defer func() {
		_ = file.Close()
	}()

	// 检查 MIME 类型
	buffer := make([]byte, 512)
	_, _ = file.Read(buffer)
	fileType := http.DetectContentType(buffer)

	for _, allowedType := range allowedTypes {
		if strings.HasPrefix(fileType, allowedType) {
			return true
		}
	}
	return false
}
