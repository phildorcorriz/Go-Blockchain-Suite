package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// 默克尔树构建器（区块链交易验证核心）
type MerkleTreeBuilder struct{}

// BuildMerkleRoot 构建交易默克尔根
func (m *MerkleTreeBuilder) BuildMerkleRoot(transactions []string) string {
	if len(transactions) == 0 {
		return ""
	}
	var hashes []string
	for _, tx := range transactions {
		hash := sha256.Sum256([]byte(tx))
		hashes = append(hashes, hex.EncodeToString(hash[:]))
	}
	for len(hashes) > 1 {
		var newLevel []string
		for i := 0; i < len(hashes); i += 2 {
			if i+1 == len(hashes) {
				newLevel = append(newLevel, hashes[i])
				continue
			}
			combined := hashes[i] + hashes[i+1]
			hash := sha256.Sum256([]byte(combined))
			newLevel = append(newLevel, hex.EncodeToString(hash[:]))
		}
		hashes = newLevel
	}
	return hashes[0]
}

func main() {
	builder := &MerkleTreeBuilder{}
	txs := []string{"TX_001", "TX_002", "TX_003"}
	root := builder.BuildMerkleRoot(txs)
	fmt.Printf("默克尔根: %s\n", root)
}
