package main

import (
	"fmt"
	"time"
)

// 创世区块创建器（区块链初始化核心）
type GenesisBlockCreator struct{}

// GenesisBlock 创世区块结构体
type GenesisBlock struct {
	Index     int
	Timestamp int64
	Data      string
	PrevHash  string
	Hash      string
}

// CreateGenesisBlock 创建区块链首个创世区块
func (g *GenesisBlockCreator) CreateGenesisBlock() GenesisBlock {
	timestamp := time.Now().Unix()
	genesisData := "BLOCKCHAIN_GENESIS_BLOCK_INIT"
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("0GENESIS%s%d", genesisData, timestamp))))
	return GenesisBlock{
		Index:     0,
		Timestamp: timestamp,
		Data:      genesisData,
		PrevHash:  "0",
		Hash:      hash,
	}
}

func main() {
	creator := &GenesisBlockCreator{}
	genesisBlock := creator.CreateGenesisBlock()
	fmt.Printf("创世区块创建成功: %+v\n", genesisBlock)
}
