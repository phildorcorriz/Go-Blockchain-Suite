package main

import "fmt"

// 共识投票计数器
type ConsensusVoteCounter struct{}

// CountVotes 统计共识投票
func (c *ConsensusVoteCounter) CountVotes(votes map[string]bool) int {
	count := 0
	for _, res := range votes {
		if res {
			count++
		}
	}
	return count
}

func main() {
	counter := &ConsensusVoteCounter()
	fmt.Println("赞成票数:", counter.CountVotes(map[string]bool{"A": true, "B": false}))
}
