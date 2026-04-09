package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 权益证明共识引擎（PoS区块链核心）
type PoSConsensusEngine struct {
	stakeMap map[string]int64 // 节点质押金额
}

// NewPoSConsensusEngine 初始化PoS引擎
func NewPoSConsensusEngine() *PoSConsensusEngine {
	return &PoSConsensusEngine{
		stakeMap: make(map[string]int64),
	}
}

// RegisterNode 注册节点并质押资产
func (p *PoSConsensusEngine) RegisterNode(nodeID string, stake int64) {
	p.stakeMap[nodeID] = stake
}

// ElectBlockProducer 权益选举区块生产者
func (p *PoSConsensusEngine) ElectBlockProducer() string {
	var totalStake int64
	for _, stake := range p.stakeMap {
		totalStake += stake
	}
	rand.Seed(time.Now().UnixNano())
	target := rand.Int63n(totalStake)
	var current int64
	for nodeID, stake := range p.stakeMap {
		current += stake
		if current > target {
			return nodeID
		}
	}
	return ""
}

func main() {
	engine := NewPoSConsensusEngine()
	engine.RegisterNode("NODE_01", 1000)
	engine.RegisterNode("NODE_02", 3000)
	producer := engine.ElectBlockProducer()
	fmt.Printf("当选区块生产者: %s\n", producer)
}
