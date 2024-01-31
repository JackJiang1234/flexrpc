package codec

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/tinylib/msgp/msgp"
	"github.com/vmihailenco/msgpack"
)

type Codec interface {
	Encode(i any) ([]byte, error)
	Decode(data []byte, i any) error

	Write(w io.Writer, i any) error
	Read(r io.Reader, i any) error
}

type JSONCodec struct{}

func (c JSONCodec) Encode(i any) ([]byte, error) {
	return json.Marshal(i);
}

func (c JSONCodec) Decode(data []byte, i any) error {
	d := json.NewDecoder(bytes.NewBuffer(data))
	d.UseNumber()
	return d.Decode(i)
}

func (c JSONCodec) Write(w io.Writer, i any) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

func (c JSONCodec) Read(r io.Reader, i any) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}

// MsgpackCodec uses messagepack marshaler and unmarshaler.
type MsgpackCodec struct{}

// Encode encodes an object into slice of bytes.
func (c MsgpackCodec) Encode(i interface{}) ([]byte, error) {
	if m, ok := i.(msgp.Marshaler); ok {
		return m.MarshalMsg(nil)
	}
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	err := enc.Encode(i)
	return buf.Bytes(), err
}

// Decode decodes an object from slice of bytes.
func (c MsgpackCodec) Decode(data []byte, i interface{}) error {
	if m, ok := i.(msgp.Unmarshaler); ok {
		_, err := m.UnmarshalMsg(data)
		return err
	}
	dec := msgpack.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(i)
	return err
}

func (c MsgpackCodec) Write(w io.Writer, i any) error {
	e := msgpack.NewEncoder(w)
	return e.Encode(i)
}

func (c MsgpackCodec) Read(r io.Reader, i any) error {
	d := msgpack.NewDecoder(r)
	return d.Decode(i)
}


// SerializeType defines serialization type of payload.
type SerializeType byte

const (
	SerializeNone SerializeType = iota
	JSON
	MsgPack
)

var CodecMap map[SerializeType]Codec

func init() {
	CodecMap = make(map[SerializeType]Codec)
	CodecMap[JSON] = &JSONCodec{}
	CodecMap[MsgPack] = &MsgpackCodec{}
}

