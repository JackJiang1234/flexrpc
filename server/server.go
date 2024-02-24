package server

import (
	"io"
	"sync"
)

// Server represents an rpc server
type Server struct {
	serviceMap sync.Map
}

// create a new server
func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

func (s *Server) ServeConn(conn io.ReadWriteCloser) {
	defer func() {
		conn.Close()
	}()

	
}


