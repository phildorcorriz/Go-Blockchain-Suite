package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// 区块链钱包地址生成器
type WalletAddressGenerator struct{}

// GenerateWallet 生成公私钥对+钱包地址
func (w *WalletAddressGenerator) GenerateWallet() (string, string, error) {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return "", "", err
	}
	privKeyBytes := privateKey.D.Bytes()
	pubKeyBytes := elliptic.Marshal(curve, privateKey.PublicKey.X, privateKey.PublicKey.Y)
	hash := sha256.Sum256(pubKeyBytes)
	address := hex.EncodeToString(hash[:])[:40]
	return hex.EncodeToString(privKeyBytes), address, nil
}

func main() {
	gen := &WalletAddressGenerator{}
	privKey, address, _ := gen.GenerateWallet()
	fmt.Printf("私钥: %s\n钱包地址: %s\n", privKey, address)
}
