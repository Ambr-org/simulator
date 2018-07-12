///Copyright (c) 2018 Ambr project
///Written by KimiKan
///Distributed under the MIT software license, see the accompanying
///file COPYING or http://www.opensource.org/licenses/mit-license.php.

package protocol

import (
	"errors"

	"github.com/boltdb/bolt"
)

const (
	FILENAME       = "default.db"
	bucketUnits    = "UNITS"
	bucketAccounts = "ACCOUNTS"
)

type DB struct {
	Context *bolt.DB
}

func CloseDB(db *DB) {
	if db == nil {
		return
	}
	db.Close()
}

func DefaultDB() (*DB, error) {
	return NewDB(FILENAME)
}

func NewDB(name string) (*DB, error) {
	db, e := bolt.Open(name, 0600, nil)
	if e != nil {
		return nil, e
	}
	p := &DB{
		Context: db,
	}
	return p, nil
}

func (p *DB) Close() {
	if p.Context != nil {
		p.Context.Close()
	}
}

func (p *DB) GetUnit(key HashKeyType) (*Unit, error) {
	var content []byte
	p.Context.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketUnits))
		if b != nil {
			content = b.Get(key.Bytes())
		}
		return nil
	})

	if content != nil {
		u, e := UnMarshalUnit(content)
		if e == nil {
			return u, nil
		}

		return nil, e
	}

	return nil, errors.New("key not exists")
}

func (p *DB) SaveUnit(u *Unit) error {
	bs, err := u.Marshal()
	if err != nil {
		return err
	}

	return p.Context.Update(func(tx *bolt.Tx) error {
		b, e := tx.CreateBucketIfNotExists([]byte(bucketUnits))
		if e != nil {
			return e
		}
		return b.Put(u.HashKey.Bytes(), bs)
	})
}

func (p *DB) GetAccount(address string) (*Account, error) {
	var content []byte
	p.Context.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketAccounts))
		if b != nil {
			content = b.Get([]byte(address))
		}
		return nil
	})

	if content != nil {
		u, e := UnMarshalAccount(content)
		if e == nil {
			return u, nil
		}
		return nil, e
	}
	return nil, errors.New("key not exists")
}

func (p *DB) SaveAccount(u *Account) error {
	bs, err := u.Marshal()
	if err != nil {
		return err
	}
	return p.Context.Update(func(tx *bolt.Tx) error {
		b, e := tx.CreateBucketIfNotExists([]byte(bucketAccounts))
		if e != nil {
			return e
		}
		return b.Put([]byte(u.Address), bs)
	})
}
