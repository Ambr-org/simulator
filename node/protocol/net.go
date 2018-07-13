///Copyright (c) 2018 Ambr project
///Written by KimiKan
///Distributed under the MIT software license, see the accompanying
///file COPYING or http://www.opensource.org/licenses/mit-license.php.

package protocol


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
