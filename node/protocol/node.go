package protocol

import (
	"errors"
	"sync"
)

type Node struct {
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

func (p *Node) GetAccount(address string) *Account {
	p.RLock()
	defer p.RUnlock()
	acc, ok := p.Accounts[address]
	if ok {
		return acc
	}
	return nil
}
