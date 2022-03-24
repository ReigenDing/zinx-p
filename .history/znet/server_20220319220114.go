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

	go func() {
		for {

			conn, err := server.Accept()
			if err != nil {
				fmt.Printf("err => %s", err)
				continue
			}
			buffer := make([]byte, 100)
			_, err = conn.Read(buffer)
			if err != nil {
				fmt.Printf("server[%s] read with error => %s", s.name, err)
				continue
			}
			fmt.Printf("message from client => %s", buffer)
			_, err = conn.Write([]byte(fmt.Sprintf("%s%s", buffer, " <=> got it!\n")))
			if err != nil {
				fmt.Printf("server[%s] write with error => %s", s.name, err)
			}
		}
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
