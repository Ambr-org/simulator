///Copyright (c) 2018 Ambr project
///Written by KimiKan
///Distributed under the MIT software license, see the accompanying
///file COPYING or http://www.opensource.org/licenses/mit-license.php.

package protocol

import (
	"bytes"
	"encoding/gob"
	"errors"
	"log"
	"time"
)

const (
	UnitSend = 1
	UnitRecv = 2
)

type HashKeyType struct {
	Value string
}

var (
	DefaultHashKey = HashKeyType{Value: ""}
)

func (p *HashKeyType) IsNullOrEmpty() bool {
	return len(p.Value) <= 0
}

func (p *HashKeyType) Bytes() []byte {
	return []byte(p.Value)
}

type Unit struct {
	Previous HashKeyType

	//for another unit. s=>r, r=>s
	OtherUnit      HashKeyType
	UnitType       int32
	AccountBalance int64
	//why place next in unit, due to
	//hash calc method disclude the next
	Next    HashKeyType
	HashKey HashKeyType
	//self owner
	Creator *Key

	//for sign->verify
	Pair      *Pair
	TimeStamp time.Time `json:"TimeStamp"`
	//Creator.verify(getData(Payload), Pair)
}

func NewUnit(creator *Key, previous HashKeyType, ut int32, balance int64) (*Unit, error) {
	u := &Unit{
		UnitType:       ut,
		AccountBalance: balance,
		Creator:        creator,
		Previous:       previous,
		TimeStamp:      time.Now(),
	}
	err := u.UpdateHash()
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (p *Unit) GetAmount(db *DB) (int64, error) {
	if db == nil {
		return 0, errors.New("invalid parameter")
	}
	if p.Previous.IsNullOrEmpty() {
		return p.AccountBalance, nil
	}
	u, err := db.GetUnit(p.Previous)
	if err != nil {
		return 0, err
	}
	return p.AccountBalance - u.AccountBalance, nil
}

func (p *Unit) GetOtherUnit(db *DB) (*Unit, error) {
	if db == nil {
		return nil, errors.New("invalid parameter")
	}
	u, e := db.GetUnit(p.OtherUnit)
	if e != nil {
		return nil, e
	}
	return u, e
}

func (p *Unit) GetOtherAccount(db *DB) (*Account, error) {
	u, e := p.GetOtherUnit(db)
	if e != nil {
		return nil, e
	}
	pub := FromKey(u.Creator)
	if pub == nil {
		return nil, errors.New("invalid creator")
	}
	addr, e2 := pub.ToAddress()
	if e2 != nil {
		return nil, e2
	}
	return db.GetAccount(addr)
}

func (p *Unit) Equal(u *Unit) bool {
	if u == nil {
		return false
	}

	return u.UnitType == p.UnitType &&
		u.AccountBalance == p.AccountBalance &&
		u.Next == p.Next &&
		u.HashKey == p.HashKey &&
		u.Previous == p.Previous &&
		u.OtherUnit == p.OtherUnit
}

func (p *Unit) NeedSignedBuffer() ([]byte, error) {
	bufs, e := GetBytes(p.Creator, p.Previous, p.UnitType, p.AccountBalance, p.TimeStamp)
	if e != nil {
		return nil, e
	}
	return ArrayToBuf(bufs)
}

func (p *Unit) NeedSignedBuffer2() ([]byte, error) {
	bufs, e := GetBytes(p.Creator, p.Previous, p.UnitType, p.AccountBalance, p.TimeStamp, p.OtherUnit)
	if e != nil {
		return nil, e
	}
	return ArrayToBuf(bufs)
}

func (p *Unit) SignSend(key *PrivateKey) error {
	buf, err := p.NeedSignedBuffer()
	if err != nil {
		return err
	}
	return p.Sign(key, buf)
}

func (p *Unit) SignRecv(key *PrivateKey) error {
	buf, err := p.NeedSignedBuffer2()
	if err != nil {
		return err
	}
	return p.Sign(key, buf)
}

func (p *Unit) Sign(key *PrivateKey, buf []byte) error {
	if key == nil || buf == nil {
		return errors.New("error key")
	}
	data, err := Marshal(buf)
	if err != nil {
		return err
	}
	pair, e2 := key.Sign(data)
	if e2 != nil {
		return e2
	}
	p.Pair = pair
	return nil
}

func (p *Unit) Update(db *DB) error {
	if db == nil {
		return errors.New("invalid parameter")
	}
	return db.SaveUnit(p)
}

func (p *Unit) UpdateHash() error {
	hash, err := GetObjectHash(p.Creator, p.Previous, p.UnitType, p.AccountBalance, p.TimeStamp)
	if err != nil {
		return err
	}
	s := string(hash)
	p.HashKey = HashKeyType{Value: s}
	return nil
}

func (p *Unit) Marshal() ([]byte, error) {
	return Marshal(p)
}

func UnMarshalUnit(b []byte) (*Unit, error) {
	u := &Unit{}
	var buf = bytes.Buffer{}
	buf.Write(b)
	// Create a decoder and receive a value.
	dec := gob.NewDecoder(&buf)
	err := dec.Decode(u)
	if err != nil {
		log.Fatal("decode:", err)
		return nil, err
	}

	return u, nil
}
