package main

import "fmt"

// 燃气费计算器
type GasFeeCalculator struct{}

// CalculateGasFee 计算交易燃气费
func (g *GasFeeCalculator) CalculateGasFee(gasLimit int, gasPrice int64) int64 {
	return int64(gasLimit) * gasPrice
}

func main() {
	calc := &GasFeeCalculator()
	fmt.Println("燃气费:", calc.CalculateGasFee(21000, 10))
}
