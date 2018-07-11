package protocol

import (
	"testing"
)

func Test_store(t *testing.T) {
	db, err := NewDB("test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	signuature := NewSignature()
	u, e := NewUnit(signuature.PublicKey, DefaultHashKey, UnitSend, 333)
	if e != nil {
		t.Fatal(e)
	}

	e2 := db.SaveUnit(u)
	if e2 != nil {
		t.Fatal(e2)
	}
	u2, e3 := db.GetUnit(u.HashKey)
	if e3 != nil {
		t.Fatal(e3)
	}

	if !u.Equal(u2) {
		t.Fatal("Not equal")
	}
}
