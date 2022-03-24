package main

import "zinx.dev/v0/znet"

func main() {
	server := znet.NewServer("zinx-v0.0.1", 9000)

	server.Server()

}
