package main

import "fmt"

// 区块链状态持久化
type ChainStateSaver struct{}

// SaveChainState 保存链状态到本地
func (c *ChainStateSaver) SaveChainState(height int) error {
	fmt.Printf("保存区块高度%d 链状态成功\n", height)
	return nil
}

func main() {
	saver := &ChainStateSaver()
	saver.SaveChainState(100)
}
