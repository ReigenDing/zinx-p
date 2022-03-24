package znet

import (
	"fmt"
	"net"

	"zinx.dev/v0/ziface"
)

type Server struct {
	name      string
	IPVersion string
	IP        string
	Port      int
}

func CallbackToClient(conn *net.TCPConn, data []byte, cin int) error {
	fmt.Printf("[handle callback]")
	_, err := conn.Write(data)
	if err != nil {
		fmt.Printf("callback write with error => %s", err)
		return err
	}
	return nil
}

func (s *Server) Start() {
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Printf("[%s] start with err => %s", s.name, err)
		return
	}
	server, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Printf("server with error => %s", err)
	}
	var cid uint = 0

	go func() {

		fmt.Printf("server wainting for connection")
		conn, err := server.AcceptTCP()
		if err != nil {
			fmt.Printf("err => %s", err)
			return
		}
		c := NewConnection(conn, cid, CallbackToClient)
		go c.Start()
		cid++
	}()

}

func (s *Server) Stop() {}

func (s *Server) Server() {
	s.Start()
	select {}
}

func NewServer(name string) ziface.IServer {
	return &Server{
		name:      name,
		IPVersion: "tcp4",
		Port:      9000,
	}

}
