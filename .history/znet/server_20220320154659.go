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

func callback(conn *net.Conn) {}

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
	cid := 0

	go func() {

		conn, err := server.Accept()
		if err != nil {
			fmt.Printf("err => %s", err)
			return
		}
		c := NewConnection(*conn, cid, callback)
		c.Start()
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
