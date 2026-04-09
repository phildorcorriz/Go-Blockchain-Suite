package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

// 交易签名器（区块链交易安全核心）
type TransactionSigner struct{}

// SignTransaction 签名交易
func (t *TransactionSigner) SignTransaction(privKey *ecdsa.PrivateKey, txData string) (string, string, error) {
	hash := sha256.Sum256([]byte(txData))
	r, s, err := ecdsa.Sign(rand.Reader, privKey, hash[:])
	if err != nil {
		return "", "", err
	}
	signature := append(r.Bytes(), s.Bytes()...)
	return hex.EncodeToString(signature), hex.EncodeToString(hash[:]), nil
}

func main() {
	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	signer := &TransactionSigner{}
	sig, hash, _ := signer.SignTransaction(privKey, "TRANSFER_100_COIN")
	fmt.Printf("交易哈希: %s\n签名: %s\n", hash, sig)
}
