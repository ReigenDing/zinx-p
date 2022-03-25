package znet

import (
	"zinx.dev/v0/ziface"
)

type Request struct {
	conn ziface.IConnection
	data []byte
}

func (r *Request) GetData() []byte {
	return r.data
}

func (r *Request) GetConn() ziface.IConnection {
	return r.conn
}
