package message

import (
	"bytes"
	"testing"
	"xrpc/transport/codec"
)

func TestNewHeader(t *testing.T) {
	h := NewHeader()
	if !h.CheckMagicNumber() {
		t.Errorf("new header has error. %v", h)
	}
}

func TestReadHeader(t *testing.T) {
	var buf bytes.Buffer
	buf.WriteByte(magicNumber)
	buf.WriteByte(byte(codec.JSON))

	h := NewHeader()
	h.Read(&buf)

	if !h.CheckMagicNumber() {
		t.Error("read header check magic number failed.")
	}
	if h.GetSerializeType() != codec.JSON {
		t.Errorf("read header check serialize type failed. expect: %v, actual: %v", codec.JSON, h.GetSerializeType())
	}
}

func TestWriteHeader(t *testing.T) {
	var buf bytes.Buffer
	h := NewHeader()
	h.SerializeType(codec.JSON)
	n, _ := h.Write(&buf)
	if (n != len(h)){
		t.Errorf("write head len is incorrected. expect: %v, actual %v", len(h), n)
	}
	
}
