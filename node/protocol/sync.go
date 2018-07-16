///Copyright (c) 2018 Ambr project
///Written by KimiKan
///Distributed under the MIT software license, see the accompanying
///file COPYING or http://www.opensource.org/licenses/mit-license.php.

package protocol

const (
	//invalid message
	Invalid = 0
	//start transfer publish
	MsgPublishSend = 1
	MsgPublishRecv = 2
	//request for more blocks
	MsgRequest = 3
	//withion a reliable network
	//why we need ack?  due to loadbalance
	//for example no ack during a time
	//we request again to another peer
	MsgRequestAck = 4
	//response for more blocks
	MsgResponse    = 5
	MsgResponseAck = 6
	//for conflict resolve usage
	MsgVote         = 7
	MsgVoteResponse = 8
	//keep alive,
	MsgHeartbeat = 1000
)

type Message struct {
	Type int32
	//
}

type Network struct {
}

func NewNetwork() *Network {

	return nil
}

func (p *Network) OnMsgReceived(m *Message) {

}

func (p *Network) Broadcast(u *Unit) error {
	return nil
}

func (p *Network) Vote(u1, u2 HashKeyType) error {
	return nil
}
