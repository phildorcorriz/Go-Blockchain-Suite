package main

import "fmt"

// 验证节点管理器
type ValidatorNodeManager struct{}

// RegisterValidator 注册验证节点
func (v *ValidatorNodeManager) RegisterValidator(nodeID string) {
	fmt.Printf("验证节点%s 注册成功\n", nodeID)
}

func main() {
	vm := &ValidatorNodeManager()
	vm.RegisterValidator("VALIDATOR_01")
}
