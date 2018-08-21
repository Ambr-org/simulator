package main

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"

	"github.com/ethereum/go-ethereum/common"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main2() {
	mdb := ethdb.NewMemDatabase()
	trieDB := trie.NewDatabase(mdb)
	tr, err := trie.New(common.Hash{}, trieDB)
	checkErr(err)
	key := []byte("abc")
	value := []byte("123")
	checkErr(tr.TryUpdate(key, value))

	key = []byte("ab")
	value = []byte("abc")
	checkErr(tr.TryUpdate(key, value))

	fmt.Println("----------")
	root, err := tr.Commit(func(leaf []byte, parent common.Hash) error {
		return nil
	})
	checkErr(err)
	tr2, err := trie.New(root, trieDB)
	checkErr(err)
	key = []byte("abcd")
	value = []byte("你好")
	checkErr(tr2.TryUpdate(key, value))
}

func Test_statedb(t *testing.T) {
	main2()
}
