package scrypt

import (
	"encoding/hex"
	"golang.org/x/crypto/scrypt"
	"gopkg.in/encoder.v1/types"
)

type Encoder struct {
	SaltLen   int // bytes to use as salt (octets)
	salt      []byte
	N         int // CPU/memory cost parameter (logN)
	R         int // block size parameter (octets)
	P         int // parallelisation parameter (positive int)
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
	var err error
	e.hashValue, err = scrypt.Key([]byte(src), e.salt, e.N, e.R, e.P, e.SaltLen)
	return e.hashValue, err
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
	encoded, err := scrypt.Key([]byte(rawData), e.salt, e.N, e.R, e.P, e.SaltLen)
	if err != nil {
		return false, err
	}
	return hash == hex.EncodeToString(encoded), nil
}

// GetSalt Returns the salt if present, otherwise nil
func (e *Encoder) GetSalt() ([]byte, error) {
	return e.salt, nil
}
