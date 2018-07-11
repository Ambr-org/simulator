package protocol

import (
	"bytes"
	"testing"
)

func Test_base58(t *testing.T) {
	src := []byte("324")
	str := Base58Encode(src)

	src2 := Base58Decode(str)

	if !bytes.Equal(src, src2) {
		t.Fatal("base58 fucking wrong")
	}
}
