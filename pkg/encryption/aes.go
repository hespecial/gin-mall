package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

var key = []byte("_HelloWorld_AnExampleMoneySecret") // AES 32bytes

// EncryptAES AES 加密
func EncryptAES(plaintext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 填充原始数据
	plaintextBytes := _PKCS7Padding([]byte(plaintext), aes.BlockSize)

	// 初始化向量 IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 加密模式
	cb := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintextBytes))
	cb.CryptBlocks(ciphertext, plaintextBytes)

	// 返回 IV 和密文
	return fmt.Sprintf("%x", append(iv, ciphertext...)), nil
}

// DecryptAES AES 解密
func DecryptAES(ciphertext string) (string, error) {
	// 解析 IV 和密文
	ivBytes, err := hex.DecodeString(ciphertext[:32])
	if err != nil {
		return "", err
	}
	iv := ivBytes
	ciphertextBytes, err := hex.DecodeString(ciphertext[32:])
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 解密模式
	cb := cipher.NewCBCDecrypter(block, iv)
	cb.CryptBlocks(ciphertextBytes, ciphertextBytes)

	// 去除填充
	plaintextBytes := _PKCS7UnPadding(ciphertextBytes)

	return string(plaintextBytes), nil
}

// PKCS7Padding 填充函数
func _PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

// PKCS7UnPadding 去除填充
func _PKCS7UnPadding(src []byte) []byte {
	length := len(src)
	if length == 0 {
		return nil
	}
	unpading := int(src[length-1])
	if unpading < 1 || unpading > length {
		return src
	}
	return src[:(length - unpading)]
}
