package znet

import "net"

type Connection struct {
	Conn     *net.TCPConn
	IsClosed bool
	ExitChan chan int
	callback func() error
}

// 启动链接方法
func (c *Connection) Start() {}

// 停止链接方法
func (c *Connection) Stop() {}

// 获取当前链接对象
func (c *Connection) GetConn() *net.TCPConn {}

// 获取当前链接的客户端的地址和端口
func (c *Connection) GetClientInfo() *net.TCPAddr {}

// 发送数据方法
func (c *Connection) Send() {}
