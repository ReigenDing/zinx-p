package znet

import (
	"fmt"
	"net"

	"zinx.dev/v0/ziface"
)

type Connection struct {
	Conn     *net.TCPConn
	ConnID   uint
	IsClosed bool
	ExitChan chan bool
	Router   ziface.IRouter
}

// 初始化链接
func NewConnection(conn *net.TCPConn, connID uint, router ziface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		IsClosed: false,
		ExitChan: make(chan bool, 1),
		Router:   router,
	}
}

func (c *Connection) StartReader() {

	defer fmt.Printf("conn %d reader is exited, remote addr => %s", c.ConnID, c.Conn.RemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Printf("conn %d read with error => %s", c.ConnID, err)
			continue
		}
		req := Request{
			conn: c,
			data: buf,
		}
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
		// if err := c.HandleAPI(c.Conn, buf, 512); err != nil {
		// 	fmt.Printf("conn %d handle message error => %s", c.ConnID, err)
		// 	break
		// }
	}
}

// 启动链接方法
func (c *Connection) Start() {
	fmt.Printf("connection %d starting", c.ConnID)
	//启动读取数据goroutine
	go c.StartReader()
	// TODO启动当前链接写逻辑
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
