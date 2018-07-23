package protocol

import (
	"errors"

	"github.com/golang/protobuf/proto"
)

func FromUnitToProto(u *Unit) (proto.Message, error) {
	if u == nil {
		return nil, errors.New("invalid parameter")
	}
	addr, err := u.Creator.ToAddress()
	if err != nil {
		return nil, err
	}
	sig, e1 := u.Pair.GetBuffer()
	if e1 != nil {
		return nil, e1
	}

	payload := &Payload{
		Balance:   u.AccountBalance,
		Creator:   addr,
		Previous:  u.Previous.Value,
		Signature: sig,
	}

	if u.UnitType == UnitSend {
		return &SendUnit{
			Payload: payload,
		}, nil
	} else if u.UnitType == UnitRecv {
		return &RecvUnit{
			Payload: payload,
			Other:   u.OtherUnit.Value,
		}, nil
	}

	return nil, errors.New("unknown unit type")
}

func FromProtoSendToUnit(p *SendUnit) (*Unit, error) {
	if p == nil {
		return nil, errors.New("invalid parameter")
	}
	pub, err := FromAddress(p.Payload.Creator)
	if err != nil {
		return nil, err
	}
	key := pub.GetKeyData()

	pair, e := UnMarshalPair(p.Payload.Signature)
	if e != nil {
		return nil, e
	}
	return NewUnit2(key, HashKeyType{
		Value: p.Payload.Previous,
	}, UnitSend, p.Payload.Balance, pair)
}

func FromProtoRecvToUnit(p *RecvUnit) (*Unit, error) {
	if p == nil {
		return nil, errors.New("invalid parameter")
	}
	pub, err := FromAddress(p.Payload.Creator)
	if err != nil {
		return nil, err
	}
	key := pub.GetKeyData()

	pair, e := UnMarshalPair(p.Payload.Signature)
	if e != nil {
		return nil, e
	}
	u, e2 := NewUnit2(key, HashKeyType{
		Value: p.Payload.Previous,
	}, UnitRecv, p.Payload.Balance, pair)
	if e2 != nil {
		return nil, e2
	}
	u.OtherUnit = HashKeyType{
		Value: p.Other,
	}
	return u, nil
}
