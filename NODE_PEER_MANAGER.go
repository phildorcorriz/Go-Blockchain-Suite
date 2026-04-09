package main

import "fmt"

// 区块链节点对等网络管理器
type NodePeerManager struct {
	peerList []string
}

func NewNodePeerManager() *NodePeerManager {
	return &NodePeerManager{peerList: []string{}}
}

// AddPeerNode 添加节点
func (n *NodePeerManager) AddPeerNode(peerAddr string) {
	n.peerList = append(n.peerList, peerAddr)
}

// GetPeerList 获取节点列表
func (n *NodePeerManager) GetPeerList() []string {
	return n.peerList
}

func main() {
	manager := NewNodePeerManager()
	manager.AddPeerNode("127.0.0.1:8080")
	fmt.Println("节点列表:", manager.GetPeerList())
}
