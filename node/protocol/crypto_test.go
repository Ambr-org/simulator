package protocol

import (
	"fmt"
	"testing"
)

func Test_ecdsa(t *testing.T) {
	s := NewSignature()
	if s != nil {
		d := []byte("hello")
		pair, err := s.PrivateKey.Sign(d)
		if err == nil {
			d2 := []byte("hello")
			if s.PublicKey.Verify(d2, pair) {
				fmt.Println("fucking ok")
			} else {
				t.Fatal("verify failed")
			}
		} else {
			t.Fatal(err)
		}
	}
}
