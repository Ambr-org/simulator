///Copyright (c) 2018 Ambr project
///Written by KimiKan
///Distributed under the MIT software license, see the accompanying
///file COPYING or http://www.opensource.org/licenses/mit-license.php.

package main

import (
	"fmt"
	"log"
	"nanosimulator/protocol"
)

func main2() {

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

func main() {
	db, err := protocol.NewDB("test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	signuature := protocol.NewSignature()
	u, e := protocol.NewUnit(signuature.PublicKey.GetKeyData(), protocol.DefaultHashKey, protocol.UnitSend, 333)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(u.TimeStamp)
	e2 := db.SaveUnit(u)
	if e2 != nil {
		log.Fatal(e2)
	}

	u2, e3 := db.GetUnit(u.HashKey)
	if e3 != nil {
		log.Fatal(e3)
	}

	fmt.Println(u2.TimeStamp)
	fmt.Println(u, u2)
	if !u.Equal(u2) {
		log.Fatal("Not equal")
	}
}
