package random

import (
	"fmt"
	"time"
)

func GenerateOrderNumber() string {
	now := time.Now()
	timestamp := now.Format("20060102150405")        // 以年月日时分秒的形式作为时间戳
	randomPart := fmt.Sprintf("%04d", r.Intn(10000)) // 生成四位随机数
	orderNumber := fmt.Sprintf("%s%s", timestamp, randomPart)
	return orderNumber
}
