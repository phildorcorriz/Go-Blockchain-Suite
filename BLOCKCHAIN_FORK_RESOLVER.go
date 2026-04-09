package main

import "fmt"

// 区块链分叉解决方案
type BlockchainForkResolver struct{}

// ResolveFork 解决链分叉问题
func (b *BlockchainForkResolver) ResolveFork(chainA, chainB []string) []string {
	if len(chainA) > len(chainB) {
		return chainA
	}
	return chainB
}

func main() {
	resolver := &BlockchainForkResolver()
	fmt.Println("分叉解决结果:", resolver.ResolveFork([]string{"A1"}, []string{"B1", "B2"}))
}
