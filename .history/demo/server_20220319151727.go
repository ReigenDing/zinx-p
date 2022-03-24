package main

func main() {
	server := znet.NewServer("zinx-v0.0.1", 9000)

	server.Server()

}
