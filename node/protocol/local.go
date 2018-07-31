package protocol

import (
	"errors"
	"log"
	"sync"

	"github.com/golang/protobuf/proto"
)

type localEndpoint struct {
	index      int32
	network    *LocalNetwork
	mailbox    chan *MsgHeader
	dispatcher Dispatcher
}

//local endpoint
func newLocalEndpoint(net *LocalNetwork, index int32, disp Dispatcher) *localEndpoint {

	ep := &localEndpoint{
		network:    net,
		index:      index,
		mailbox:    make(chan *MsgHeader),
		dispatcher: disp,
	}
	go func() {
		m := <-ep.mailbox
		if m != nil {
			e := ep.OnMsgReceived(m.sender, m.message)
			if e != nil {
				log.Fatal(e)
			}
		}
	}()
	return ep
}

func (p *localEndpoint) equals(other *localEndpoint) bool {
	if other == nil {
		return false
	}
	return p.index == other.index
}

//publish a message to network
func (p *localEndpoint) Publish(m proto.Message) error {
	return p.network.Publish(m, p)
}

//msg dispacher
func (p *localEndpoint) OnMsgReceived(sender string, m proto.Message) error {
	if m == nil {
		return errors.New("invalid parameter")
	}

	switch t := m.(type) {
	case *SendUnit:
		return p.dispatcher.OnSendUnitArrived(sender, t)
	case *RecvUnit:
		return p.dispatcher.OnRecvUnitArrived(sender, t)
	case *VoteRequest:
		return p.dispatcher.OnVoteRequest(sender, t)
	case *VoteResponse:
		return p.dispatcher.OnVoteResponse(sender, t)
	case *HeartbeatRequest:
		return p.dispatcher.OnHeartbeatRequest(sender, t)
	case *HeartbeatResponse:
		return p.dispatcher.OnHeartbeatResponse(sender, t)
	case *ReplicationRequest:
		return p.dispatcher.OnReplicationRequest(sender, t)
	case *ReplicationResponse:
		return p.dispatcher.OnReplicationResponse(sender, t)
	default:
		return errors.New("Unexpected type")
	}
}

type LocalNetwork struct {
	sync.RWMutex
	nodes map[int32]*localEndpoint
}

func newLocalNetwork() *LocalNetwork {
	return &LocalNetwork{
		nodes: make(map[int32]*localEndpoint),
	}
}

func (p *LocalNetwork) newEndpoint(index int32, disp Dispatcher) *localEndpoint {
	p.Lock()
	defer p.Unlock()
	ep := newLocalEndpoint(p, index, disp)
	p.nodes[index] = ep

	return ep
}

func (p *LocalNetwork) getEndpoint(index int32) *localEndpoint {
	p.RLock()
	defer p.RUnlock()
	ep, ok := p.nodes[index]
	if ok {
		return ep
	}

	return nil
}

//publish a message to network
func (p *LocalNetwork) Publish(m proto.Message, sender *localEndpoint) error {
	if sender == nil || m == nil {
		return errors.New("invalid parameter")
	}

	p.RLock()
	defer p.RUnlock()
	for k, v := range p.nodes {
		if k != sender.index {
			//mailbox
			v.mailbox <- &MsgHeader{
				sender:  string(sender.index),
				message: m,
			}
		}
	}
	return nil
}
