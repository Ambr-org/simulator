package protocol

import (
	"bytes"
	"fmt"
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

func Test_base58_decode(t *testing.T) {
	src := "434&^%^"
	str := Base58Decode(src)
	fmt.Println(str)
	if str != nil {
		t.Fatal("base58 fucking wrong")
	}
}
