package protocol

import (
	"testing"
)

func Test_cfg(t *testing.T) {
	c := GetConfig_mock("../cfg.txt")
	if c == nil {
		t.Fatal("nil config")
	}

	if c.P2p.Port <= 0 {
		t.Fatal("error content")
	}
}
