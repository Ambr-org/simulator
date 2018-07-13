package protocol

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_base64(t *testing.T) {
	src := []byte("324")
	str := Base64Encode(src)

	src2, err := Base64Decode(str)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(src, src2) {
		t.Fatal("base64 fucking wrong")
	}
}

func Test_base64_decode(t *testing.T) {
	src := "434&^%^"
	_, err := Base64Decode(src)

	if err == nil {
		t.Fatal("base64 fucking wrong")
	} else {
		fmt.Println(err)
	}
}
