package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetConn() *net.TCPConn
	GetClientInfo() net.TCPAddr
	Send()
}

type HandleFunc func(*net.TCPConn, []byte, int) error
