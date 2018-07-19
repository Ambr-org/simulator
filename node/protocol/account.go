///Copyright (c) 2018 Ambr project
///Written by KimiKan
///Distributed under the MIT software license, see the accompanying
///file COPYING or http://www.opensource.org/licenses/mit-license.php.

package protocol

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
)

type Account struct {
	Owner     *Key
	FirstUnit HashKeyType
	LastUnit  HashKeyType

	//should be key for account struct
	Address     string
	DB          *DB
	Transporter Transporter

	//pending units need to be confirm
	//by user
	PendingRecvUnits []HashKeyType
}

func LoadAccount(db *DB, hash HashKeyType) (*Account, error) {
	return nil, nil
}

func NewAccount(db *DB, key *Key, tr Transporter) *Account {
	if db == nil || key == nil || tr == nil {
		return nil
	}
	addr, err := key.ToAddress()
	if err != nil {
		return nil
	}

	return &Account{
		Owner:       key,
		Transporter: tr,
		Address:     addr,
	}
}

//VerifyUnit defines how to verify the unit
//u represents a recv unit
func (p *Account) VerifyRecvUnit(u *Unit) error {
	if u == nil {
		return errors.New("invalid unit")
	}
	if u.Pair == nil {
		return errors.New("empty signature")
	}

	pub := FromKey(u.Creator)
	if pub == nil {
		return errors.New("invalid creator")
	}
	buf, e := u.NeedSignedBuffer2()
	if e != nil {
		return e
	}
	if !pub.Verify(buf, u.Pair) {
		return errors.New("Verify failed, signature not correct")
	}
	u2, e2 := p.DB.GetUnit(u.OtherUnit)
	if e2 != nil {
		return e2
	}
	if u2.OtherUnit.IsNullOrEmpty() {
		fmt.Println("fucking, conflict, vote needed")
		//return p.Transporter.Vote(u.HashKey, u2.OtherUnit)
		//TODO: vote message
		return p.Transporter.Publish(nil)
	}
	amount1, ex1 := u.GetAmount(p.DB)
	if ex1 != nil {
		return ex1
	}

	amount2, ex2 := u2.GetAmount(p.DB)
	if ex2 != nil {
		return ex2
	}

	if amount2 != amount1 {
		//not quite equals
		//abandon it
		return errors.New("not equal, error state")
	}

	u2.OtherUnit = u.HashKey
	return u2.Update(p.DB)
}

//VerifyUnit defines how to verify the unit
//u represents a recv unit
func (p *Account) VerifySendUnit(u *Unit) error {
	if u == nil {
		return errors.New("invalid unit")
	}
	if u.Pair == nil {
		return errors.New("empty signature")
	}

	pub := FromKey(u.Creator)
	if pub == nil {
		return errors.New("invalid creator")
	}
	buf, e := u.NeedSignedBuffer()
	if e != nil {
		return e
	}
	if !pub.Verify(buf, u.Pair) {
		return errors.New("Verify failed, signature not correct")
	}
	if !u.OtherUnit.IsNullOrEmpty() {
		return errors.New("error state, decline it")
	}

	//every unit should have a valid previous
	if u.Previous.IsNullOrEmpty() && p.FirstUnit.IsNullOrEmpty() {
		return errors.New("error previous unit")
	}

	//only verify the latest one
	//every node should maintain the correct state
	account, e := p.GetAccountBalance()
	if account <= 0 {
		return errors.New("will not recv any spend at all")
	}

	//next, money of couse.
	if u.AccountBalance == account {
		return errors.New("0 spend not allowed")
	}
	if u.AccountBalance < 0 {
		return errors.New("account corrupted")
	}

	return nil
}

func (p *Account) Update() error {
	if p.DB == nil {
		return errors.New("invalid db context")
	}
	return p.DB.SaveAccount(p)
}

func (p *Account) Marshal() ([]byte, error) {
	return Marshal(p)
}

func (p *Account) GetAccountBalance() (int64, error) {
	if p.DB == nil {
		return 0, errors.New("invalid db context")
	}
	if p.LastUnit.IsNullOrEmpty() {
		return 0, nil
	}
	u, err := p.DB.GetUnit(p.LastUnit)
	if err != nil {
		return 0, err
	}
	return u.AccountBalance, nil
}

//generate unit
//sign it, then publish it to peers
func (p *Account) StartTransfer(from *PrivateKey, to *PublicKey, amount int64) error {
	if from == nil || to == nil || amount <= 0 {
		return errors.New("error parameters")
	}
	balance, e := p.GetAccountBalance()
	if e != nil {
		return e
	}

	if balance < amount {
		return errors.New("balance not enough")
	}

	u, err := NewUnit(p.Owner, p.LastUnit, UnitSend, balance-amount)
	if err != nil {
		return err
	}

	if p.FirstUnit.IsNullOrEmpty() {
		p.FirstUnit = u.HashKey
	}

	if p.LastUnit.IsNullOrEmpty() {
		p.LastUnit = u.HashKey
	} else {
		u2, e2 := p.DB.GetUnit(p.LastUnit)
		if e2 != nil {
			return e2
		}
		u2.Next = u.HashKey
		u.Previous = u2.HashKey
		p.LastUnit = u2.HashKey
		//save u, u2, account
		//should within a transaction
		//but demo , you know
		e3 := p.DB.SaveUnit(u2)
		if e3 != nil {
			return e3
		}
	}

	ex := u.SignSend(from)
	if ex != nil {
		return ex
	}
	//save u, account
	e4 := p.DB.SaveUnit(u)
	if e4 != nil {
		return e4
	}

	e5 := p.DB.SaveAccount(p)
	if e5 != nil {
		return e5
	}

	//start transfer
	//TODO: createdunit, send?
	return p.Transporter.Publish(FromUnitToProto(u))
}

func (p *Account) ConfirmTransfer(u *Unit) error {
	e := p.VerifyRecvUnit(u)
	if e != nil {
		return e
	}
	if p.FirstUnit.IsNullOrEmpty() {
		p.FirstUnit = u.HashKey
	}

	if p.LastUnit.IsNullOrEmpty() {
		p.LastUnit = u.HashKey
	} else {
		u2, e2 := p.DB.GetUnit(p.LastUnit)
		if e2 != nil {
			return e2
		}
		u2.Next = u.HashKey
		u.Previous = u2.HashKey
		p.LastUnit = u2.HashKey
		//save u, u2, account
		//should within a transaction
		//but demo , you know
		e3 := p.DB.SaveUnit(u2)
		if e3 != nil {
			return e3
		}
	}

	//save u, account
	e4 := p.DB.SaveUnit(u)
	if e4 != nil {
		return e4
	}

	e5 := p.DB.SaveAccount(p)
	if e5 != nil {
		return e5
	}

	//confirm
	return p.Transporter.Publish(FromUnitToProto(u))
}

func UnMarshalAccount(b []byte) (*Account, error) {
	u := &Account{}
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
