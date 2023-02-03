package hmac

import (
	"crypto/hmac"
	"encoding/hex"
	"hash"
)

type Encoder struct {
	HashFunc  func() hash.Hash // hash func default is sha256.New
	key       string           // is the secret key
	hashValue []byte
}

// Hash Generate and return a hash value in []byte format
func (e *Encoder) Hash(src string) ([]byte, error) {
	if e.hashValue != nil {
		return e.hashValue, nil
	}
	h := hmac.New(e.HashFunc, []byte(e.key))
	h.Write([]byte(src))
	e.hashValue = h.Sum(nil)
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
	decoded, err := hex.DecodeString(hash)
	if err != nil {
		return false, err
	}
	data, err := e.Hash(rawData)
	if err != nil {
		return false, err
	}
	return hmac.Equal(data, decoded), nil
}

// GetSalt Returns the salt if present, otherwise nil
func (e *Encoder) GetSalt() ([]byte, error) {
	return nil, nil
}
