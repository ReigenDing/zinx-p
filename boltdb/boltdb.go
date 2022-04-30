package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

// type struct1 struct {
// 	a1 uint32
// 	a2 uint32
// 	a3 uint32
// 	a4 uint32
// 	a5 uint32
// }

func main() {

	// st1 := struct1{a1: 1 << 31, a2: 1 << 31, a3: 1 << 31, a4: 1 << 31, a5: 1 << 31}
	// fmt.Printf("%v", st1)
	// fmt.Printf("%d\n", unsafe.Sizeof(st1))

	// fmt.Printf("%d\n", int(9>>1))

	// a1 := [5]int{1, 2, 6, 9, 44}
	// fmt.Printf("%v\n", a1[4:])
	// index := sort.Search(5, func(i int) bool { fmt.Println(i); return a1[i] > 6 })
	// fmt.Println("=============================")
	// fmt.Println(index)

	// fmt.Println(syscall.Getpagesize())
	// fmt.Println(0x12 == 0x1)

	db, err := bolt.Open("./my.db", 0600, nil)
	fmt.Printf("after db open opetation\n")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		fmt.Printf("before create bucket opetation\n")
		bucket, err := tx.CreateBucketIfNotExists([]byte("user4"))
		if err != nil {
			log.Fatalf("CreateBucketIfNotExists err %s", err.Error())
			return err
		}
		b5, err := bucket.CreateBucketIfNotExists([]byte("user5"))
		bb5 := bucket.Bucket([]byte("user5"))
		fmt.Printf("bucket user5 in user4 => %v, %v", b5, bb5)
		// 放入东西
		fmt.Printf("before put key in bucket opetation\n")
		if err = b5.Put([]byte("hello5"), []byte("world5")); err != nil {
			log.Fatalf("bucket Put err: %s", err.Error())
			return err
		}
		fmt.Printf("after put key in bucket opetation\n")
		return nil
	})

	fmt.Printf("after db put opetation\n")
	if err != nil {
		log.Fatalf("db.Update err: %s", err.Error())
	}
	// db中读取数据
	err = db.View(func(tx *bolt.Tx) error {
		// 找到柜子
		log.Printf("the get bucket: %s", "before")
		bucket := tx.Bucket([]byte("user"))
		fmt.Println(bucket)
		// 找东西的人
		// log.Printf("the get val: %s", "before")
		// val := bucket.Get([]byte("hello"))
		// log.Printf("the get val: %s", val)
		// val = bucket.Get([]byte("hello2"))
		// log.Fatalf("the get val21: %s", val)
		return nil
	})
	if err != nil {
		log.Fatalf("db.View err: %s", err.Error())
	}
}
