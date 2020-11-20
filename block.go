package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

//1. 定义结构
type Block struct {
	//1. 版本号
	Version uint64
	//2. 前区块哈希
	PrevHash []byte
	//3. Merkel根(就是一个哈希值，先不管)
	MerkelRoot []byte
	//4. 时间戳
	TimeStamp uint64
	//5. 难度值
	Diffculty uint64
	//6. 随机数，挖矿要找的数
	Nonce uint64
	//7. 当前区块哈希 正常比特币区块中没有当前区块的哈希，为了方便简化//TODO
	Hash []byte
	//8. 数据
	Data []byte
}
//实现一个辅助函数，将uint64转成[]byte
func Uint64ToByte(num uint64) []byte {
	//TODO
	var buffer bytes.Buffer

	//通过重新二进制编码转换
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}
//2. 创建区块
func NewBlock(data string, preBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   preBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Diffculty:  0,
		Nonce:      0,
		Hash:       []byte{},
		Data:       []byte(data),
	}
	//设置哈希值
	block.SetHash()

	return &block
}

//3. 生成哈希
func (block *Block) SetHash() {
	// 1. 拼装数据
	/*
	blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo, block.PrevHash...)
	blockInfo = append(blockInfo, block.MerkelRoot...)
	blockInfo = append(blockInfo, block.Data...)
	blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Diffculty)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
	*/
	tmp:= [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		block.Data,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Diffculty),
		Uint64ToByte(block.Nonce),
	}
	//将二维切片数组链接起来，返回一个一维切片数组
	blockInfo := bytes.Join(tmp, []byte{})
	// 2. sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
