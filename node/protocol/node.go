package protocol

import (
	"errors"
	"sync"
)

type Node struct {
	Dispatcher
	sync.RWMutex
	Accounts    map[string]*Account
	DB          *DB
	Transporter Transporter
}

func NewNode(db *DB, trans Transporter) *Node {
	if db == nil || trans == nil {
		return nil
	}

	node := &Node{
		Accounts:    make(map[string]*Account),
		DB:          db,
		Transporter: trans,
	}

	return node
}

func (p *Node) AccountExists(addr string) bool {
	if _, ok := p.Accounts[addr]; ok {
		return true
	}

	return false
}

func (p *Node) GetOrCreateAccount(addr string) (*Account, error) {
	if acc, ok := p.Accounts[addr]; ok {
		return acc, nil
	}
	account, err := p.NewAccountWithAddress(addr)
	if err != nil {
		return nil, err
	}

	err1 := account.Update()
	if err1 != nil {
		return nil, err1
	}
	err2 := p.AddAccount(account)
	if err2 != nil {
		return nil, err2
	}
	return account, nil
}

func (p *Node) AddAccount(account *Account) error {
	if account == nil {
		return errors.New("invalid parameter")
	}

	if len(account.Address) == 0 {
		return errors.New("invalid address")
	}
	p.Lock()
	defer p.Unlock()
	if _, ok := p.Accounts[account.Address]; ok {
		return errors.New("account already exists")
	}
	p.Accounts[account.Address] = account
	return nil
}

func (p *Node) NewAccountWithAddress(address string) (*Account, error) {
	if !IsValidAddress(address) {
		return nil, errors.New("invalid address")
	}
	pub, err := FromAddress(address)
	if err != nil {
		return nil, err
	}
	key := pub.GetKeyData()
	return NewAccount(p.DB, key, p.Transporter), nil
}

func (p *Node) GetAccount(address string) *Account {
	p.RLock()
	defer p.RUnlock()
	acc, ok := p.Accounts[address]
	if ok {
		return acc
	}
	return nil
}

//func transfer try to
func (p *Node) Transfer(sender string, target string, amount int64) error {
	return nil
}

//send unit received
func (p *Node) OnSendUnitArrived(sender string, m *SendUnit) error {
	if m == nil {
		return errors.New("invlalid paramenter")
	}

	return nil
}

//sender is ip or something address else
func (p *Node) OnRecvUnitArrived(sender string, m *RecvUnit) error {
	if m == nil {
		return errors.New("invlalid paramenter")
	}
	if !IsValidAddress(m.Payload.Creator) {
		return errors.New("invalid creator")
	}

	account, err := p.GetOrCreateAccount(m.Payload.Creator)
	if err != nil {
		return err
	}
	u, e := FromProtoRecvToUnit(m)
	if e != nil {
		return e
	}

	return account.ConfirmTransfer(u)
}

//vote. while conflict
func (p *Node) OnVoteRequest(sender string, m *VoteRequest) error {
	return nil
}

func (p *Node) OnVoteResponse(sender string, m *VoteResponse) error {
	return nil
}

//heartbeat to keep peer alived
//if not provided it's okay
//libary maintained
func (p *Node) OnHeartbeatRequest(sender string, m *HeartbeatRequest) error {
	return nil
}

func (p *Node) OnHeartbeatResponse(sender string, m *HeartbeatResponse) error {
	return nil
}

//for replication used
//request for lost nodes
//it should be carefully designed
func (p *Node) OnReplicationRequest(sender string, m *ReplicationRequest) error {
	return nil
}

func (p *Node) OnReplicationResponse(sender string, m *ReplicationResponse) error {
	return nil
}
