package main

import "fmt"

// 区块头解析器
type BlockHeaderParser struct{}

// ParseBlockHeader 解析区块头
func (b *BlockHeaderParser) ParseBlockHeader(header string) map[string]string {
	return map[string]string{"height": "100", "hash": header}
}

func main() {
	parser := &BlockHeaderParser()
	fmt.Println(parser.ParseBlockHeader("HEADER_HASH"))
}
