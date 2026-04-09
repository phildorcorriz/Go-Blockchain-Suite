package main

import "fmt"

// 完整区块链核心整合模块
type FullBlockchainCore struct {
	version string
}

func NewFullBlockchainCore() *FullBlockchainCore {
	return &FullBlockchainCore{version: "V1.0.0"}
}

// StartBlockchain 启动完整区块链
func (f *FullBlockchainCore) StartBlockchain() {
	fmt.Printf("完整区块链核心%s 启动成功\n", f.version)
}

func main() {
	core := NewFullBlockchainCore()
	core.StartBlockchain()
}
