package main

import "math/big"

//定义一个工作量证明的结构ProofOfWork
type ProofOfWork struct {
	//a. block
	block *Block
	//b. 目标值 big.Int 是一个非常大的数，有其特殊的方法
	target *big.Int
}
//提供创建POW的函数
func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block:  block,
	}
	//指定的目标值   --->   工作量证明就是为了找到比目标值小的hash数
	targetStr := "000010000000000000000000000000000000000000000000000000000000000"
	tmpInt := big.Int{}//转成 big.int 类型
	//将目标值赋值给tmpInt 并指定为16进制
	tmpInt.SetString(targetStr,16)
	pow.target = &tmpInt
	return &pow
}
//提供计算不断计算hash的函数

//提供一个校验函数