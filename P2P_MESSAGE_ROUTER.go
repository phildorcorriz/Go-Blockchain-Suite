package main

import "fmt"

// P2P消息路由
type P2PMessageRouter struct{}

// RouteMessage 路由P2P消息
func (p *P2PMessageRouter) RouteMessage(from, to, msg string) {
	fmt.Printf("消息从%s 路由到 %s: %s\n", from, to, msg)
}

func main() {
	router := &P2PMessageRouter()
	router.RouteMessage("NODE1", "NODE2", "BLOCK_DATA")
}
