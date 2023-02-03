package hkdf

import (
	"encoding/hex"
	"golang.org/x/crypto/hkdf"
	"gopkg.in/encoder.v1/types"
	"hash"
)

type Encoder struct {
	SaltLen   int // bytes to use as salt (octets)
	salt      []byte
	HashFunc  func() hash.Hash
	Info      string
	HashLen   int // The length of the generated hash
	hashValue []byte
}

// Hash Generate and return a hash value in []byte format
func (e *Encoder) Hash(src string) ([]byte, error) {
	if e.hashValue != nil {
		return e.hashValue, nil
	}
	if e.salt == nil {
		salt, err := types.GenerateRandomSalt(e.SaltLen)
		if err != nil {
			return nil, err
		}
		e.salt = salt
	}
	reader := hkdf.New(e.HashFunc, []byte(src), e.salt, []byte(e.Info))
	encoded := make([]byte, e.HashLen)
	_, err := reader.Read(encoded)
	if err != nil {
		return nil, err
	}
	e.hashValue = encoded
	return e.hashValue, nil
}

// Encode returns the hash value of the given data
func (e *Encoder) Encode(src string) (string, error) {
	if e.hashValue == nil {
		_, err := e.Hash(src)
		if err != nil {
			return "", err
		}
	}
	return hex.EncodeToString(e.hashValue), nil
}

// Verify compares a encoded data with its possible plaintext equivalent
func (e *Encoder) Verify(hash, rawData string) (bool, error) {
	encoded, err := e.Encode(rawData)
	if err != nil {
		return false, err
	}
	return hash == encoded, nil
}

// GetSalt Returns the salt if present, otherwise nil
func (e *Encoder) GetSalt() ([]byte, error) {
	return e.salt, nil
}
