package protocol

import "github.com/golang/protobuf/proto"

type MsgHeader struct {
	sender  string
	message proto.Message
}
