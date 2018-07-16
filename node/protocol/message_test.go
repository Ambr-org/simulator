package protocol

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
)

func Test_message(t *testing.T) {
	s := &TestMessage{
		Message: "hello",
		Length:  20,
		Cnt:     32,
	}

	//protobuf
	pData, err := proto.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}

	var out TestMessage
	e1 := proto.Unmarshal(pData, &out)
	if e1 != nil {
		t.Fatal(e1)
	}
	fmt.Println(out)
}
