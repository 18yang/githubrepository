package main

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
		Hash:     []byte{},//后面进行填充//TODO
		Data:     []byte(data),
	}
	return &block
}
//3. 生成哈希

//4. 引入区块链

//5. 添加区块

//6. 重构代码


func main() {
	NewBlock("你太强了",[]byte{})
}
