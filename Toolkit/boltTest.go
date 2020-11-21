package main

import (
	"BlockChain/bolt"
	"fmt"
	"log"
)

func main() {
	//1. 打开数据库
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Panic("打开数据库失败")
	}

	//2. 找到抽屉bucket（没有就创建）
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			//没有抽屉，创建
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic("创建bucket(b1)失败")
			}
		}
		//3. 写数据
		bucket.Put([]byte("11111"),[]byte("hello"))
		bucket.Put([]byte("22222"),[]byte("world"))
		fmt.Println("创建成功")
		return nil
	})

	//4. 读数据
}
