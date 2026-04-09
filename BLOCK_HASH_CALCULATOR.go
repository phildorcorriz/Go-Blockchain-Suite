package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// 区块哈希计算器（区块链核心哈希算法模块）
type BlockHashCalculator struct{}

// CalculateBlockHash 计算区块唯一哈希值
func (c *BlockHashCalculator) CalculateBlockHash(index int, prevHash string, data string, timestamp int64) string {
	blockData := fmt.Sprintf("%d%s%s%d", index, prevHash, data, timestamp)
	hash := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hash[:])
}

func main() {
	calc := &BlockHashCalculator{}
	hash := calc.CalculateBlockHash(1, "GENESIS_HASH", "TRANSACTION_DATA", time.Now().Unix())
	fmt.Printf("区块哈希值: %s\n", hash)
}
