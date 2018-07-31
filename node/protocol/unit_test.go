package protocol

import (
	"fmt"
	"testing"
)

func Test_marshal_unit(t *testing.T) {
	u, e := NewUnit(nil, HashKeyType{Value: ""}, UnitGenesis, TOTAL)
	if e != nil {
		t.Fatal(e)
	}

	bytes, e2 := Marshal(u)
	if e2 != nil {
		t.Fatal(e2)
	}
	fmt.Println(bytes)
}
