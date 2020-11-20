package main

import (
	"fmt"
)

func main() {
	chain := NewBlockChain()

	chain.AddBlock("A向B转了50BTN")
	chain.AddBlock("A向C转了50BTN")

	for i, block := range chain.blocks {
		fmt.Printf("============当前区块高度：%d============\n", i)
		fmt.Printf("前区块的哈希值:  %x\n", block.PrevHash)
		fmt.Printf("当前区块的哈希值: %x\n", block.Hash)
		fmt.Printf("当前区块的数据:  %s\n", block.Data)
	}
}
