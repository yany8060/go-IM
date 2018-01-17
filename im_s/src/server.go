package main

import (
	"net"
)

type ServerHub struct {
	baddr *net.TCPAddr
}
type Server struct {
	ln    net.Listener
	baddr *net.TCPAddr
}

func (s *Server) handleConn(conn net.Conn) {

}
