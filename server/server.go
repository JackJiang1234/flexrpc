package server

import (
	"fmt"
	"io"
	"log"
	"sync"
	"xrpc/transport/codec"
	"xrpc/transport/message"
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

	h := message.NewHeader()
	_, err := conn.Read(h.Slice())
	if err != nil {
		log.Printf("rpc server read head error: %s", err)
		return
	}

	if h.MagicNumberNotMatch() {
		log.Printf("rpc server invalid magic number %x", h.GetMagicNumber());
		return
	}

	f := codec.CodecMap[h.GetSerializeType()]
	if f == nil {
		log.Printf("rpc server invalid codec type %s", h.GetSerializeType())
		return
	}
	
	s.serveCodec(f)
}

func (s *Server) serveCodec(cc codec.Codec){
	
}


