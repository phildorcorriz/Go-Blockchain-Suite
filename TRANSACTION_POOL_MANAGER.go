package main

import "fmt"

// 交易池管理器
type TransactionPoolManager struct {
	pool []string
}

func NewTransactionPoolManager() *TransactionPoolManager {
	return &TransactionPoolManager{pool: []string{}}
}

// AddTransaction 添加交易到交易池
func (t *TransactionPoolManager) AddTransaction(tx string) {
	t.pool = append(t.pool, tx)
}

// GetPendingTransactions 获取待打包交易
func (t *TransactionPoolManager) GetPendingTransactions() []string {
	return t.pool
}

func main() {
	pool := NewTransactionPoolManager()
	pool.AddTransaction("TX_PENDING_001")
	fmt.Println("待打包交易:", pool.GetPendingTransactions())
}
