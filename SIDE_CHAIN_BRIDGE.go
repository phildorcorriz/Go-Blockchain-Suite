package main

import "fmt"

// 侧链跨链桥
type SideChainBridge struct{}

// SyncToSideChain 同步数据到侧链
func (s *SideChainBridge) SyncToSideChain(data string) string {
	return "同步到侧链: " + data
}

func main() {
	bridge := &SideChainBridge()
	fmt.Println(bridge.SyncToSideChain("BLOCK_DATA"))
}
