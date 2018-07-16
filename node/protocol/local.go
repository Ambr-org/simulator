package protocol

import (
	"errors"
	"sync"
)

type localEndpoint struct {
	node    *Node
	index   int32
	network *LocalNetwork
}

//local endpoint
func newLocalEndpoint(net *LocalNetwork, index int32, node *Node) *localEndpoint {
	ep := &localEndpoint{
		network: net,
		index:   index,
		node:    node,
	}
	return ep
}

func (p *localEndpoint) equals(other *localEndpoint) bool {
	if other == nil {
		return false
	}
	return p.index == other.index
}

//publish a message to network
func (p *localEndpoint) Publish(m interface{}) error {
	return p.network.Publish(m, p)
}

//msg dispacher
func (p *localEndpoint) OnMsgReceived(m interface{}) error {
	if m == nil {
		return errors.New("invalid parameter")
	}

	switch t := m.(type) {
	case *CreatedUnit:
		return p.OnUnitCreated(t)
	case *VoteRequest:
		return p.OnVoteRequest(t)
	case *VoteResponse:
		return p.OnVoteResponse(t)
	case *HeartbeatRequest:
		return p.OnHeartbeatRequest(t)
	case *HeartbeatResponse:
		return p.OnHeartbeatResponse(t)
	case *ReplicationRequest:
		return p.OnReplicationRequest(t)
	case *ReplicationResponse:
		return p.OnReplicationResponse(t)
	default:
		return errors.New("Unexpected type")
	}
}

//created unit received
func (p *localEndpoint) OnUnitCreated(m *CreatedUnit) error {
	return nil
}

//vote. while conflict
func (p *localEndpoint) OnVoteRequest(m *VoteRequest) error {
	return nil
}

func (p *localEndpoint) OnVoteResponse(m *VoteResponse) error {
	return nil
}

//heartbeat to keep peer alived
//if not provided it's okay
//libary maintained
func (p *localEndpoint) OnHeartbeatRequest(m *HeartbeatRequest) error {
	return nil
}

func (p *localEndpoint) OnHeartbeatResponse(m *HeartbeatResponse) error {
	return nil
}

//for replication used
//request for lost nodes
//it should be carefully designed
func (p *localEndpoint) OnReplicationRequest(m *ReplicationRequest) error {
	return nil
}

func (p localEndpoint) OnReplicationResponse(m *ReplicationResponse) error {
	return nil
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
func (p *LocalNetwork) Publish(m interface{}, sender *localEndpoint) error {
	if sender == nil || m == nil {
		return errors.New("invalid parameter")
	}

	p.RLock()
	defer p.RUnlock()
	for k, v := range p.nodes {
		if k != sender.index {
			err := v.OnMsgReceived(m)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
