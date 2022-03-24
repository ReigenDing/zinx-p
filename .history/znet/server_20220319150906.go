package znet

import (
	"fmt"
	"net"
)

type Server struct {
	name      string
	IPVersion string
	IP        string
	Port      int
}

func NewServer(name: string, port: int) *Server {
    return &Server{
        name: name,
        IPVersion: "ipv4",
    }

}

func (s *Server) Start() {
	addr, err := net.ResolveTCPAddr(s.IPVersion, s.IP)
	if err != nil {
		fmt.Printf("[%s] start with err => %s", s.name, err)
		return
	}
	server, err := net.Listen("tcp", addr.IP.String())

	go func() {
		for {

			conn, err := server.Accept()
			if err != nil {
				fmt.Printf("err => %s", err)
				continue
			}
			buffer := make([]byte, 100)
			conn.Read(buffer)

			conn.Write(buffer)
		}
	}()

}

func (s *Server) stop() {}

func (s *Server) Server() {
	s.Start()
	select {}
}
