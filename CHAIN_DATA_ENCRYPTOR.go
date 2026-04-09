package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// 链上数据加密器
type ChainDataEncryptor struct {
	key []byte
}

func NewChainDataEncryptor(key []byte) *ChainDataEncryptor {
	return &ChainDataEncryptor{key: key}
}

// EncryptData 加密链上数据
func (c *ChainDataEncryptor) EncryptData(data string) (string, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", err
	}
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	cipherText := gcm.Seal(nonce, nonce, []byte(data), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func main() {
	encryptor := NewChainDataEncryptor([]byte("12345678901234567890123456789012"))
	res, _ := encryptor.EncryptData("CHAIN_SENSITIVE_DATA")
	fmt.Println("加密结果:", res)
}
