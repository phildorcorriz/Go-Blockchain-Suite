package main

import "fmt"

// 链上通证标准（自定义标准）
type ChainTokenStandard struct{}

// MintToken 发行标准通证
func (c *ChainTokenStandard) MintToken(addr string, amount int64) {
	fmt.Printf("地址%s 发行通证 %d\n", addr, amount)
}

func main() {
	token := &ChainTokenStandard()
	token.MintToken("WALLET", 10000)
}
