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
	response, err := client.Write([]byte("nihao,golang"))
	if err != nil {
		fmt.Printf("client write with error => %s", err)
	}
	fmt.Printf("server reponse => %s", response)
}
