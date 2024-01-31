package message

import (
	"io"
	"xrpc/transport/codec"
)

// the frist byte is magic number
// the second byte is serialize type
type Header [2]byte

type Body struct {
	Seq uint64 // sequence number chosen by client

	Error         string
	ServicePath   string
	ServiceMethod string
	Args          any
	Metadata      map[string]string
	Payload       []byte
}

const (
	magicNumber byte = 0x06
)

func (h Header) CheckMagicNumber() bool {
	return h[0] == magicNumber
}

func (h Header) GetSerializeType() codec.SerializeType {
	return codec.SerializeType(h[1])
}

func (h Header) SerializeType(t codec.SerializeType) {
	h[1] = byte(t)
}

func (h *Header) Read(r io.Reader) (n int, err error) {
	return r.Read(h[:])
}

func (h *Header) Write(w io.Writer) (n int, err error) {
	return w.Write(h[:])
}

func NewHeader() *Header {
	h := new(Header)
	h[0] = magicNumber

	return h
}
