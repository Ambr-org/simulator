///Copyright (c) 2018 Ambr project
///Written by KimiKan
///Distributed under the MIT software license, see the accompanying
///file COPYING or http://www.opensource.org/licenses/mit-license.php.

package protocol

import (
	"github.com/golang/protobuf/proto"
)

type Transporter interface {
	//publish a message to network
	Publish(m proto.Message) error
	OnMsgReceived(sender string, m proto.Message) error
}
