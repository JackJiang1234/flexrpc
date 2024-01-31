package message

import "testing"

func TestNewHeader(t *testing.T) {
	h := NewHeader();
	if (!h.CheckMagicNumber()){
		t.Errorf("new header has error. %v", h)
	}
}

