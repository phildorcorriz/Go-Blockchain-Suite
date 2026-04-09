package main

import "fmt"

// 区块数据压缩器
type BlockDataCompressor struct{}

// CompressBlock 压缩区块数据
func (b *BlockDataCompressor) CompressBlock(data string) string {
	return "COMPRESSED_" + data
}

func main() {
	comp := &BlockDataCompressor()
	fmt.Println(comp.CompressBlock("BLOCK_100_DATA"))
}
