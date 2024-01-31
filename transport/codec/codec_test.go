package codec

import (
	"bytes"
	"testing"
)

type item struct {
	Foo string
	Bar int
}

func TestJsonEncodeDecode(t *testing.T) {
	codecTest(CodecMap[JSON], t)
}

func TestJsonWriteRead(t *testing.T) {
	writeReadTest(CodecMap[JSON], t)
}

func TestMsgpackEncodeDecode(t *testing.T) {
	codecTest(CodecMap[MsgPack], t)
}

func TestMsgpackWriteRead(t *testing.T) {
	writeReadTest(CodecMap[MsgPack], t)
}

func codecTest(c Codec, t *testing.T) {
	before := item{
		Foo: "test",
		Bar: 1,
	}
	data, err := c.Encode(before)

	if err != nil {
		t.Errorf("json encode has error, %s", err)
	}

	after := item{}
	c.Decode(data, &after)
	if after != before {
		t.Errorf("json codec has error. before:%v, after: %v", before, after)
	}
}

func writeReadTest(c Codec, t *testing.T){
	var b bytes.Buffer
	before := item{
		Foo: "test",
		Bar: 1,
	}
	c.Write(&b, before)

	after := item{}
	c.Read(&b, &after)
	if after != before {
		t.Errorf("json codec has error. before:%v, after: %v", before, after)
	}
}