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


func (c *Connection) StartReader() {
	for {
		buf := []byte
		c.Conn.Read()
	}
}

// 启动链接方法
func (c *Connection) Start() {
	//启动读取数据goroutine
	go c.StartReader()
}

// 停止链接方法
func (c *Connection) Stop() {
	if c.IsClosed {
		return
	}
	c.IsClosed = true

	c.Conn.Close()

	close(c.ExitChan)
}

// 获取当前链接ID
func (c *Connection) GetConnID() uint {
	return c.ConnID
}

// 获取当前链接对象
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前链接的客户端的地址和端口
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据方法
func (c *Connection) Send(data []byte) error {
	return nil
}
