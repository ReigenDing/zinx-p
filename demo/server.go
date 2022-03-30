package main

import (
	"fmt"

	"zinx.dev/v0/ziface"
	"zinx.dev/v0/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) Handle(request ziface.IRequest) error {
	fmt.Printf("[myrouter] handle data => %s\n", request.GetData())
	if _, err := request.GetConn().GetTCPConnection().Write([]byte("pong pong pong......\n")); err != nil {
		fmt.Printf("server response with error => %s\n", err)
		return err
	}

	return nil
}

func main() {
	server := znet.NewServer("zinx-v0.0.1")
	server.AddRouter(&PingRouter{})
	server.Server()

}
