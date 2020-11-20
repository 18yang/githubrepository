package main

import (
	"crypto/sha256"
	"fmt"
)

//1. 定义结构
type Block struct {
	//1. 前区块哈希
	PrevHash []byte
	//2. 当前区块哈希
	Hash []byte
	//3. 数据
	Data []byte
}

//2. 创建区块
func NewBlock(data string , preBlockHash []byte) *Block {
	block := Block{
		PrevHash: preBlockHash,
		Hash:     []byte{},//后面进行填充
		Data:     []byte(data),
	}
	//设置哈希值
	block.SetHash()

	return &block
}
//3. 生成哈希
func (block *Block) SetHash()  {
	// 1. 拼装数据
	blockInfo := append(block.PrevHash, block.Data...)
	// 2. sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
//4. 引入区块链
type BlockChain struct {
	//定义一个区块链数组
	blocks []*Block
}
//返回一个链
func NewBlockChain() *BlockChain {
	//作为第一个创世块并加入到区块链中
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}
//定义一个创世块
func GenesisBlock() *Block {
	return NewBlock("创世块",[]byte{})
}
//5. 添加区块

//6. 重构代码


func main() {
	chain := NewBlockChain()

	for i, block := range chain.blocks  {
		fmt.Printf("============当前区块高度：%d============\n",i+1)
		fmt.Printf("前一个区块的哈希值: %x\n", block.PrevHash)
		fmt.Printf("当前区块的哈希值: %x\n", block.Hash)
		fmt.Printf("当前区块的数据:   %s\n", block.Data)
	}
}
