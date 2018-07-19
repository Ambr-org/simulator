package protocol

import (
	"errors"
	"log"
	"sync"

	"github.com/golang/protobuf/proto"
)

type localEndpoint struct {
	node    *Node
	index   int32
	network *LocalNetwork
	mailbox chan *MsgHeader
}

//local endpoint
func newLocalEndpoint(net *LocalNetwork, index int32, node *Node) *localEndpoint {
	ep := &localEndpoint{
		network: net,
		index:   index,
		node:    node,
		mailbox: make(chan *MsgHeader),
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
	case *CreatedUnit:
		return p.OnUnitCreated(sender, t)
	case *VoteRequest:
		return p.OnVoteRequest(sender, t)
	case *VoteResponse:
		return p.OnVoteResponse(sender, t)
	case *HeartbeatRequest:
		return p.OnHeartbeatRequest(sender, t)
	case *HeartbeatResponse:
		return p.OnHeartbeatResponse(sender, t)
	case *ReplicationRequest:
		return p.OnReplicationRequest(sender, t)
	case *ReplicationResponse:
		return p.OnReplicationResponse(sender, t)
	default:
		return errors.New("Unexpected type")
	}
}

//created unit received
func (p *localEndpoint) OnUnitCreated(sender string, m *CreatedUnit) error {
	return nil
}

//vote. while conflict
func (p *localEndpoint) OnVoteRequest(sender string, m *VoteRequest) error {
	return nil
}

func (p *localEndpoint) OnVoteResponse(sender string, m *VoteResponse) error {
	return nil
}

//heartbeat to keep peer alived
//if not provided it's okay
//libary maintained
func (p *localEndpoint) OnHeartbeatRequest(sender string, m *HeartbeatRequest) error {
	return nil
}

func (p *localEndpoint) OnHeartbeatResponse(sender string, m *HeartbeatResponse) error {
	return nil
}

//for replication used
//request for lost nodes
//it should be carefully designed
func (p *localEndpoint) OnReplicationRequest(sender string, m *ReplicationRequest) error {
	return nil
}

func (p localEndpoint) OnReplicationResponse(sender string, m *ReplicationResponse) error {
	return nil
}

type MsgHeader struct {
	sender  string
	message proto.Message
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

func (p *LocalNetwork) newEndpoint(index int32, node *Node) *localEndpoint {
	p.Lock()
	defer p.Unlock()
	ep := newLocalEndpoint(p, index, node)
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
