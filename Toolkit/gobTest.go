package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Person struct {
	Name string
	age uint
}

func main() {
	xiaoMing := Person{
		Name: "小明",
		age:  20,
	}

	var buffer bytes.Buffer
	//使用gob进行序列化得到字节流
	// 1. 定义一个编码器
	encoder := gob.NewEncoder(&buffer)
	//2. 使用编码器编码
	err := encoder.Encode(&xiaoMing)
	if err != nil {
		panic(err)
	}

fmt.Printf("%v\n",buffer)
	//使用gob进行反序列化
	//1. 定义一个解码器
	decoder := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))
	//2. 使用解码器解码
	var daMing Person
	err = decoder.Decode(&daMing)
	if err != nil {
		panic(err)
	}

	fmt.Printf("解码后的小明：%v\n",daMing)
}
