package main

import (
	"fmt"
	"net"
)

func main() {
	client, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Printf("client connect err => %s", err)
	}
	client.Write([]byte{"还好，golang"})
}
