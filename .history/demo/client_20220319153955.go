package main

import (
	"fmt"
	"net"
)

func main() {
	client, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Printf()
	}
}
