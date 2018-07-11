
///Copyright (c) 2018 Ambr project
///Written by KimiKan
///Distributed under the MIT software license, see the accompanying
///file COPYING or http://www.opensource.org/licenses/mit-license.php.

package main

import (
	"fmt"
	"node/protocol"
)

func main() {

	s := protocol.NewSignature()
	if s != nil {
		d := []byte("hello")
		pair, err := s.PrivateKey.Sign(d)
		if err == nil {
			d2 := []byte("hello2")
			if s.PublicKey.Verify(d2, pair) {
				fmt.Println("fucking ok")
			}
		}
	}
	fmt.Println("ok, go..")
}
