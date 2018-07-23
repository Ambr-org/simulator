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

func Test_address(t *testing.T) {
	s := NewSignature()
	if s == nil {
		t.Fatal("nil sig")
	}
	addr, ok := s.PublicKey.ToAddress()
	if ok != nil {
		t.Fatal(ok)
	}

	if !IsValidAddress(addr) {
		t.Fatal("invalid addr")
	}

	if IsValidAddress("something else") {
		t.Fatal("invalid addr")
	}
}
