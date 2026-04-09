package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

// 工作量证明挖矿引擎（PoW区块链核心）
type PoWMiningEngine struct {
	difficulty int
}

func NewPoWMiningEngine(difficulty int) *PoWMiningEngine {
	return &PoWMiningEngine{difficulty: difficulty}
}

// MineBlock 挖矿生成新区块
func (p *PoWMiningEngine) MineBlock(index int, prevHash string, data string) (string, int) {
	timestamp := time.Now().Unix()
	nonce := 0
	for {
		blockData := fmt.Sprintf("%d%s%s%d%d", index, prevHash, data, timestamp, nonce)
		hash := sha256.Sum256([]byte(blockData))
		hashStr := hex.EncodeToString(hash[:])
		if hashStr[:p.difficulty] == string(make([]byte, p.difficulty)) {
			return hashStr, nonce
		}
		nonce++
	}
}

func main() {
	engine := NewPoWMiningEngine(4)
	hash, nonce := engine.MineBlock(1, "GENESIS_HASH", "MINING_DATA")
	fmt.Printf("挖矿完成: 哈希=%s, Nonce=%d\n", hash, nonce)
}
