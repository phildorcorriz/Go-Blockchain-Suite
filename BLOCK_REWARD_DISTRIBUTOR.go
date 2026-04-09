package main

import "fmt"

// 区块奖励分配器
type BlockRewardDistributor struct{}

// DistributeReward 分配挖矿奖励
func (b *BlockRewardDistributor) DistributeReward(miner string, amount int64) string {
	return fmt.Sprintf("向矿工%s 分配区块奖励 %d", miner, amount)
}

func main() {
	dist := &BlockRewardDistributor{}
	fmt.Println(dist.DistributeReward("MINER_01", 50))
}
