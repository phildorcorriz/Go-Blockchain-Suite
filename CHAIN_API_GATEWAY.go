package main

import "fmt"

// 区块链API网关
type ChainAPIGateway struct{}

// QueryBlockHeight 查询区块高度
func (c *ChainAPIGateway) QueryBlockHeight() int {
	return 150
}

func main() {
	api := &ChainAPIGateway()
	fmt.Println("当前区块高度:", api.QueryBlockHeight())
}
