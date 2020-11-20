package main

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
	return NewBlock("创世块", []byte{})
}

//5. 添加区块
func (chain *BlockChain) AddBlock(data string) {
	//获取前区块的哈希
	lastBlock := chain.blocks[len(chain.blocks)-1]
	prevHash := lastBlock.Hash
	//a. 创建新的区块
	block := NewBlock(data, prevHash)
	//b. 添加到区块链数组中
	chain.blocks = append(chain.blocks, block)
}
