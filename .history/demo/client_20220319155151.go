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
	_, err = client.Write([]byte("nihao,golang"))
	if err != nil {
		fmt.Printf("client write with error => %s", err)
	}
	buff := make([]byte, 512)
	_, err = client.Read(buff)
	if err != nil {
		fmt.Printf("client read with error => %s", err)
	}
	fmt.Printf("server reponse => %s", buff)

}
