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
	defer db.Close()
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
		bucket.Put([]byte("11111"), []byte("hello"))
		bucket.Put([]byte("22222"), []byte("world"))

		return nil
	})

	//4. 读数据
	db.View(func(tx *bolt.Tx) error {
		//1. 找到抽屉，没有的话直接退出
		bucket := tx.Bucket([]byte("b1"))
		if err != nil {
			log.Panic("bucket b1 不应该为空，请检查！")
		}
		v1 := bucket.Get([]byte("11111"))
		v2 := bucket.Get([]byte("22222"))
		fmt.Printf("%s\n,%s\n", v1, v2)
		return nil
	})
}
