package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnID() uint
	RemoteAddr() net.TCPAddr
	Send([]byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
