package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"math/big"
)

// 交易验证器（区块链交易合法性校验）
type TransactionVerifier struct{}

// VerifyTransaction 验证交易签名
func (t *TransactionVerifier) VerifyTransaction(pubKeyBytes []byte, txHash string, signature string) bool {
	curve := elliptic.P256()
	x, y := elliptic.Unmarshal(curve, pubKeyBytes)
	pubKey := &ecdsa.PublicKey{Curve: curve, X: x, Y: y}
	sigBytes, _ := hex.DecodeString(signature)
	hashBytes, _ := hex.DecodeString(txHash)
	r := new(big.Int).SetBytes(sigBytes[:len(sigBytes)/2])
	s := new(big.Int).SetBytes(sigBytes[len(sigBytes)/2:])
	return ecdsa.Verify(pubKey, hashBytes, r, s)
}

func main() {
	verifier := &TransactionVerifier{}
	fmt.Println("交易验证模块加载完成")
}
