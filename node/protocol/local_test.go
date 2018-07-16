package protocol

import (
	"fmt"
	"log"
	"os"
	"testing"
)

const (
	kDBFile    = "users%d.db"
	kNodeCount = 10
	kUserCount = 10
)

type user struct {
	signature *Signature
	account   *Account
}

func newUser(db *DB, trans Transporter) *user {
	sig := NewSignature()
	acc := NewAccount(db, sig.PublicKey.GetKeyData(), trans)
	return &user{
		signature: sig,
		account:   acc,
	}
}

type Context struct {
	network *LocalNetwork
	nodes   []*Node
	users   []*user
}

func newContext() *Context {
	nodes := []*Node{}
	users := []*user{}
	network := newLocalNetwork()
	if network == nil {
		return nil
	}

	ctx := &Context{
		nodes:   nodes,
		network: network,
		users:   users,
	}
	ctx.init()
	return ctx
}

func (p *Context) newNode(index int32) *Node {
	db, err := NewDB(fmt.Sprintf(kDBFile, index))
	if err != nil {
		fmt.Println("new db error: ", err)
		return nil
	}

	endpoint := p.network.newEndpoint(index, nil)
	node := NewNode(db, endpoint)
	if node == nil {
		return nil
	}

	endpoint.node = node
	return node
}

func (p *Context) init() {
	for i := 0; i < kNodeCount; i++ {
		node := p.newNode(int32(i))
		if node == nil {
			log.Fatal("fucking node create failed")
		}
		p.nodes = append(p.nodes, node)
	}

	for i := 0; i < kUserCount; i++ {
		user := newUser(p.nodes[i].DB, p.network.newEndpoint(int32(i), p.nodes[i]))
		if user == nil {
			log.Fatal("fucking user create failed")
		}
		p.users = append(p.users, user)
	}
}

func (p *Context) close() {
	for _, v := range p.nodes {
		v.DB.Close()
	}
}

func clean_dbs() {
	for i := 0; i < kNodeCount; i++ {
		clean_db(fmt.Sprintf(kDBFile, i))
	}
}

func clean_db(dbfile string) {
	_, err := os.Stat(dbfile)
	if err == nil {
		e := os.Remove(dbfile)
		if e != nil {
			fmt.Println(e)
		}
	} else {
		if os.IsExist(err) {
			e := os.Remove(dbfile)
			if e != nil {
				fmt.Println(e)
			}
		}
	} //endif
}

func tear_down() {
	clean_dbs()
}

func test_setup() *Context {
	clean_dbs()
	ctx := newContext()
	/*
		u, e := NewUnit(ctx.users[0].account.Owner, DefaultHashKey, UnitRecv, 100000)
		if e != nil {
			fmt.Println(e)
		}

		fmt.Println(u)*/
	return ctx
}

func Test_transfer(t *testing.T) {
	ctx := test_setup()

	fmt.Println(ctx)

	ctx.close()
	tear_down()
	//todo: transfer
}
