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

	//NewUnit()
	return nil, nil
}

func FromProtoRecvToUnit(p *RecvUnit) (*Unit, error) {
	return nil, nil
}
