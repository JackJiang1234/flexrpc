package transport

import "xrpc/transport/message"

type Request struct {
	Msg *message.Message
} 

type Response struct {

}

type Handler interface {
    Serve(req *Request, res *Response)
}

type MessageHandle func (*message.Message)

type Server interface {
	Open() error
	Close() error

	Serve(MessageHandle) error
}