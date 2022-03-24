package znet

import "net"

type Connection struct {
	Conn *net.TCPConn

	IsClosed bool
	ExitChan chan int
	callback func() error
}
