///Copyright (c) 2018 Ambr project
///Written by KimiKan
///Distributed under the MIT software license, see the accompanying
///file COPYING or http://www.opensource.org/licenses/mit-license.php.

package protocol

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/gob"
	"errors"
	"log"
	"math/big"
	"reflect"
)

//pair struct represents
//a sign output
type Pair struct {
	R *big.Int
	S *big.Int
}

func UnMarshalPair(keyBuf []byte) (*Pair, error) {
	u := &Pair{}
	var buf = bytes.Buffer{}
	buf.Write(keyBuf)
	// Create a decoder and receive a value.
	dec := gob.NewDecoder(&buf)
	err := dec.Decode(u)
	if err != nil {
		log.Fatal("decode pair:", err)
		return nil, err
	}

	return u, nil
}

func (p *Pair) GetBuffer() ([]byte, error) {
	buf, e := Marshal(p)
	if e != nil {
		return nil, nil
	}

	return buf, nil
}

type Key struct {
	X *big.Int
	Y *big.Int
}

func UnMarshalKey(keyBuf []byte) (*Key, error) {
	u := &Key{}
	var buf = bytes.Buffer{}
	buf.Write(keyBuf)
	// Create a decoder and receive a value.
	dec := gob.NewDecoder(&buf)
	err := dec.Decode(u)
	if err != nil {
		log.Fatal("decode key:", err)
		return nil, err
	}

	return u, nil
}

func (p *Key) IntoPublicKey() *ecdsa.PublicKey {
	return &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     p.X,
		Y:     p.Y,
	}
}

func (p *Key) GetBuffer() ([]byte, error) {
	buf, e := Marshal(p)
	if e != nil {
		return nil, nil
	}

	return buf, nil
}

//address translate
//reduce p256
func (p *Key) ToAddress() (string, error) {
	buf, e := p.GetBuffer()
	if e != nil {
		return "", nil
	}
	return Base58Encode(buf), nil
}

type PublicKey struct {
	Compare
	ecdsa.PublicKey
}

func FromKey(k *Key) *PublicKey {
	p := k.IntoPublicKey()
	return &PublicKey{
		PublicKey: *p,
	}
}

//ambr address-> publickey
func FromAddress(addr string) (*PublicKey, error) {
	buf := Base58Decode(addr)
	if buf == nil {
		return nil, errors.New("invalid base58 character")
	}

	key, e := UnMarshalKey(buf)
	if e != nil {
		return nil, e
	}
	return FromKey(key), nil
}

func (p *PublicKey) Equals(o *PublicKey) bool {
	if o == nil {
		return false
	}
	return reflect.DeepEqual(p, o)
}

func (p *PublicKey) Verify(data []byte, pair *Pair) bool {
	return ecdsa.Verify(&p.PublicKey, data, pair.R, pair.S)
}

func (p *PublicKey) GetKeyData() *Key {
	return &Key{
		X: p.X,
		Y: p.Y,
	}
}

//address translate
//reduce p256
func (p *PublicKey) ToAddress() (string, error) {
	k := p.GetKeyData()
	buf, e := k.GetBuffer()
	if e != nil {
		return "", nil
	}
	return Base58Encode(buf), nil
}

type PrivateKey struct {
	Compare
	ecdsa.PrivateKey
}

type Signature struct {
	PrivateKey *PrivateKey
	PublicKey  *PublicKey
}

func NewSignature() *Signature {
	//see http://golang.org/pkg/crypto/elliptic/#P256
	pubkeyCurve := elliptic.P256()
	// this generates a public & private key pair
	//use rand seed to decrease the crash chance
	privKey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)

	if err != nil {
		return nil
	}

	return &Signature{
		PrivateKey: &PrivateKey{PrivateKey: *privKey},
		PublicKey:  &PublicKey{PublicKey: privKey.PublicKey},
	}
}

func (p *PrivateKey) Equals(o *PrivateKey) bool {
	if o == nil {
		return false
	}
	return reflect.DeepEqual(p, o)
}

func (p *PrivateKey) Sign(data []byte) (*Pair, error) {
	r, s, e := ecdsa.Sign(rand.Reader, &p.PrivateKey, data)
	return &Pair{
		R: r,
		S: s,
	}, e
}
