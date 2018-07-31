///Copyright (c) 2018 Ambr project
///Written by KimiKan
///Distributed under the MIT software license, see the accompanying
///file COPYING or http://www.opensource.org/licenses/mit-license.php.

package protocol

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"reflect"
	"time"
)

func Marshal(o interface{}) ([]byte, error) {
	var buf bytes.Buffer
	// Create an encoder and send a value.
	enc := gob.NewEncoder(&buf)
	gob.Register(Key{})
	gob.Register(HashKeyType{})
	gob.Register(time.Time{})
	err := enc.Encode(o)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func GetRandHash() ([]byte, error) {
	b := make([]byte, 1000)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return GetSHA256Hash([][]byte{b}), nil
}

func GetSHA256Hash(datas [][]byte) []byte {

	sha := sha256.New()
	for _, data := range datas {
		if data == nil {
			continue
		}
		sha.Write(data)
	}

	return sha.Sum(nil)
}

func ArrayToBuf(datas [][]byte) ([]byte, error) {
	if datas == nil {
		return nil, errors.New("error parameter")
	}
	result := []byte{}
	for _, data := range datas {
		if data == nil {
			continue
		}
		result = append(result, data...)
	}

	if len(result) <= 0 {
		return nil, errors.New("empty inputs")
	}
	return result, nil
}

func GetBytes(datas ...interface{}) ([][]byte, error) {
	if datas == nil {
		return nil, errors.New("error object")
	}

	l := [][]byte{}
	for _, data := range datas {
		t := reflect.TypeOf(data)

		if t.Kind() == reflect.Ptr && reflect.ValueOf(data).IsNil() {
			continue
		}
		buf, e := Marshal(data)
		if e != nil {
			return nil, e
		}
		//fmt.Println(datas, data, buf)
		l = append(l, buf)
	}
	return l, nil
}

func GetObjectHash(datas ...interface{}) ([]byte, error) {
	l, err := GetBytes(datas...)
	if err != nil {
		return nil, err
	}

	hash := GetSHA256Hash(l)
	return hash, nil
}
