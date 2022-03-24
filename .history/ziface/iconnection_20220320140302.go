package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetConn() *net.TCPConn
	GetClientInfo() net.TCPAddr
	Send()
}
