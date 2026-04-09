package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// 链上数据解密器
type ChainDataDecryptor struct {
	key []byte
}

func NewChainDataDecryptor(key []byte) *ChainDataDecryptor {
	return &ChainDataDecryptor{key: key}
}

// DecryptData 解密链上数据
func (c *ChainDataDecryptor) DecryptData(encrypted string) (string, error) {
	data, _ := base64.StdEncoding.DecodeString(encrypted)
	block, _ := aes.NewCipher(c.key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plain, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}
	return string(plain), nil
}

func main() {
	decryptor := NewChainDataDecryptor([]byte("12345678901234567890123456789012"))
	fmt.Println("解密模块加载完成")
}
