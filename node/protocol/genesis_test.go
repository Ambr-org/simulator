package protocol

import (
	"fmt"
	"testing"
)

func Test_genesis_sig(t *testing.T) {
	sig := NewSignature()
	privStr, e1 := sig.PrivateKey.ToString()
	if e1 != nil {
		t.Fatal(e1)
	}

	pubStr, e2 := sig.PublicKey.ToAddress()
	if e2 != nil {
		t.Fatal(e1)
	}

	t.Fatal("++", privStr, "++--", pubStr, "--")
}

func Test_genesis(t *testing.T) {
	u, err := GetGenesisUnit()
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}

	fmt.Println(u)
	//t.Fatal(u)
}
