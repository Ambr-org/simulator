package protocol

import (
	"testing"
)

func Test_getHash(t *testing.T) {
	bytes, err := GetObjectHash(nil, "test")
	if err != nil {
		t.Fatal(err)
	}

	if bytes == nil {
		t.Fatal(bytes)
	}
}

func Test_get_bytes(t *testing.T) {
	bytes, err := GetBytes(nil, "test")
	if err != nil {
		t.Fatal(err, bytes)
	}
}

/*
//should never use nil pointer for only parameter
func Test_gethash_nil(t *testing.T) {
	_, err := GetObjectHash(nil)
	if err == nil {
		t.Fatal(err)
	}
}
*/
