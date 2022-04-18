package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"
)

func g1(ch1, ch2 chan int) {
	for i := 1; i < 100; i++ {
		<-ch1
		fmt.Printf("g1 is %d\n", i)
		ch2 <- i
	}
}

func g2(ch1, ch2 chan int) {
	for i := 1; i < 100; i++ {
		<-ch2
		fmt.Printf("g2 is %d\n", i)
		ch1 <- i
	}
}

const defaultMaxFileSize = 1 << 30        // 假设文件最大为 1G
const defaultMemMapSize = 128 * (1 << 20) // 假设映射的内存大小为 128M

type Demo struct {
	file    *os.File
	data    *[defaultMaxFileSize]byte
	dataRef []byte
}

func _assert(condition bool, msg string, v ...interface{}) {
	if !condition {
		panic(fmt.Sprintf(msg, v...))
	}
}

func (demo *Demo) mmap() {
	b, err := syscall.Mmap(int(demo.file.Fd()), 0, defaultMemMapSize, syscall.PROT_WRITE|syscall.PROT_READ, syscall.MAP_SHARED)
	_assert(err == nil, "failed to mmap", err)
	demo.dataRef = b
	demo.data = (*[defaultMaxFileSize]byte)(unsafe.Pointer(&b[0]))
}

func (demo *Demo) grow(size int64) {
	if info, _ := demo.file.Stat(); info.Size() >= size {
		return
	}
	_assert(demo.file.Truncate(size) == nil, "failed to truncate")
}

func (demo *Demo) munmap() {
	_assert(syscall.Munmap(demo.dataRef) == nil, "failed to munmap")
	demo.data = nil
	demo.dataRef = nil
}

func main() {
	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// defer close(ch1)
	// defer close(ch2)
	// fmt.Printf("channel init success!!\n")
	// go g1(ch1, ch2)
	// fmt.Printf("started g1\n")
	// go g2(ch1, ch2)
	// fmt.Printf("started g2\n")
	// ch1 <- 0
	// fmt.Printf("write 0 into ch1\n")
	// time.Sleep(time.Second * 6)

	fd, err := os.OpenFile("./test.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	defer fd.Close()
	buf := []byte("hello world")
	fmt.Printf("%s", buf)
	// if _, err := fd.Write(buf); err != nil {
	// 	log.Fatal(err)
	// }
	if _, err := fd.WriteAt(buf, 0); err != nil {
		log.Fatal(err)
	}

	_ = os.Remove("tmp.txt")
	f, _ := os.OpenFile("tmp.txt", os.O_CREATE|os.O_RDWR, 0644)
	demo := &Demo{file: f}
	demo.grow(1)
	demo.mmap()
	// defer demo.munmap()

	msg := "hello geektutu!"

	demo.grow(int64(len(msg) * 2))
	for i, v := range msg {
		demo.data[2*i] = byte(v)
		demo.data[2*i+1] = byte(' ')
	}
}
