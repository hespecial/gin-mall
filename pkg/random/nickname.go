package random

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const length = 6

var r = rand.New(rand.NewSource(time.Now().Unix()))

// GenerateNickname 生成一个指定长度的随机昵称
func GenerateNickname() string {
	nickname := make([]byte, length)
	for i := range nickname {
		nickname[i] = charset[r.Intn(len(charset))]
	}
	return string(nickname)
}
