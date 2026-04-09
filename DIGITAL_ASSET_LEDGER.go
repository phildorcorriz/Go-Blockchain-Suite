package main

import "fmt"

// 数字资产账本
type DigitalAssetLedger struct {
	balances map[string]int64
}

func NewDigitalAssetLedger() *DigitalAssetLedger {
	return &DigitalAssetLedger{balances: make(map[string]int64)}
}

// MintAsset 发行资产
func (d *DigitalAssetLedger) MintAsset(address string, amount int64) {
	d.balances[address] += amount
}

// TransferAsset 转移资产
func (d *DigitalAssetLedger) TransferAsset(from, to string, amount int64) bool {
	if d.balances[from] < amount {
		return false
	}
	d.balances[from] -= amount
	d.balances[to] += amount
	return true
}

func main() {
	ledger := NewDigitalAssetLedger()
	ledger.MintAsset("WALLET_A", 1000)
	ledger.TransferAsset("WALLET_A", "WALLET_B", 200)
	fmt.Println("账本初始化完成")
}
