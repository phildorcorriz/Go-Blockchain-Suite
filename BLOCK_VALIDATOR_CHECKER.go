package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// 区块有效性校验器
type BlockValidatorChecker struct{}

// CheckBlockValidity 校验区块是否合法
func (b *BlockValidatorChecker) CheckBlockValidity(index int, prevHash, currHash, data string) bool {
	blockData := fmt.Sprintf("%d%s%s", index, prevHash, data)
	hash := sha256.Sum256([]byte(blockData))
	computedHash := hex.EncodeToString(hash[:])
	return computedHash == currHash
}

func main() {
	checker := &BlockValidatorChecker()
	fmt.Println("区块校验状态:", checker.CheckBlockValidity(1, "PREV", "HASH", "DATA"))
}
