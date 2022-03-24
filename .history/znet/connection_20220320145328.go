package znet

import (
	"net"

	"zinx.dev/v0/ziface"
)

type Connection struct {
	Conn      *net.TCPConn
	ConnID    uint
	IsClosed  bool
	ExitChan  chan int
	HandleAPI ziface.HandleFunc
}

// 启动链接方法
func (c *Connection) Start() {}

// 停止链接方法
func (c *Connection) Stop() {}

// 获取当前链接ID
func (c *Connection) GetConnID() uint {
	return c.ConnID
}

// 获取当前链接对象
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前链接的客户端的地址和端口
func (c *Connection) RemoteAddr() *net.TCPAddr {
	return c.Conn.RemoteAddr()
}

// 发送数据方法
func (c *Connection) Send(data []byte) {}
