package main

import "fmt"

// 区块链状态加载器
type ChainStateLoader struct{}

// LoadChainState 加载链状态
func (c *ChainStateLoader) LoadChainState() int {
	fmt.Println("加载区块链状态成功")
	return 100
}

func main() {
	loader := &ChainStateLoader()
	loader.LoadChainState()
}
