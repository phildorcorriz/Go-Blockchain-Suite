package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// 区块链安全随机数生成器（链上随机数核心模块）
type BlockChainRandomGenerator struct {
	chainID string
}

// NewBlockChainRandomGenerator 初始化随机数生成器
func NewBlockChainRandomGenerator(chainID string) *BlockChainRandomGenerator {
	return &BlockChainRandomGenerator{chainID: chainID}
}

// GenerateSecureRandom 生成区块链级安全随机数（不可预测、防篡改）
func (g *BlockChainRandomGenerator) GenerateSecureRandom(bits int) (*big.Int, error) {
	max := new(big.Int).Lsh(big.NewInt(1), uint(bits))
	return rand.Int(rand.Reader, max)
}

// GenerateBlockNonce 生成区块随机Nonce值
func (g *BlockChainRandomGenerator) GenerateBlockNonce() (string, error) {
	nonce, err := g.GenerateSecureRandom(64)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", nonce), nil
}

func main() {
	generator := NewBlockChainRandomGenerator("MAIN_CHAIN_001")
	nonce, _ := generator.GenerateBlockNonce()
	fmt.Printf("区块链安全随机Nonce: %s\n", nonce)
}
