///Copyright (c) 2018 Ambr project
///Written by KimiKan
///Distributed under the MIT software license, see the accompanying
///file COPYING or http://www.opensource.org/licenses/mit-license.php.

package protocol

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	"reflect"
)

//pair struct represents
//a sign output
type Pair struct {
	R *big.Int
	S *big.Int
}

type Key struct {
	X *big.Int
	Y *big.Int
}

func (p *Key) IntoPublicKey() *ecdsa.PublicKey {
	return &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     p.X,
		Y:     p.Y,
	}
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

func (p *PublicKey) Equals(o *PublicKey) bool {
	if o == nil {
		return false
	}
	return reflect.DeepEqual(p, o)
}

func (p *PublicKey) Verify(data []byte, pair *Pair) bool {
	return ecdsa.Verify(&p.PublicKey, data, pair.R, pair.S)
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
