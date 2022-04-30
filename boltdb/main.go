package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {

	file, err := os.OpenFile("my.db", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stat, err := os.Stat("my.db")
	if err != nil {
		panic(err)
	}
	b, err := syscall.Mmap(int(file.Fd()), 0, int(stat.Size()), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		panic(err)
	}
	defer syscall.Munmap(b)
	// b[0] = byte(1)
	// b[1034] = byte(255)
	// b[1035] = byte(255)
	// b[1036] = byte(255)
	fmt.Println(b[:4096])
	fmt.Println(b[4096 : 4096*2])
	fmt.Println(b[4096*2 : 4096*3])
	fmt.Println(b[4096*3 : 4096*4])
	fmt.Println(b[4096*4 : 4096*5])
	fmt.Println(b[4096*5 : 4096*6])
	fmt.Println(b[4096*6 : 4096*7])

}
