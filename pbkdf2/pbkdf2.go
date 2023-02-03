package pbkdf2

import (
	"crypto/rand"
	"encoding/hex"
	"hash"

	"golang.org/x/crypto/pbkdf2"
)

const (
	DefaultIterations = 10000
	alphabet          = "./ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

type Encoder struct {
	salt       []byte
	SaltLen    int
	Iterations int
	KeyLen     int
	HashFunc   func() hash.Hash
	hashValue  []byte
}

// Hash Generate and return a hash value in []byte format
func (e *Encoder) Hash(src string) ([]byte, error) {
	if e.hashValue != nil {
		return e.hashValue, nil
	}
	e.generateSalt(e.SaltLen)
	e.hashValue = pbkdf2.Key([]byte(src), e.salt, e.Iterations, e.KeyLen, e.HashFunc)
	return e.hashValue, nil
}

func (e *Encoder) generateSalt(length int) {
	if e.salt == nil {
		salt := make([]byte, length)
		rand.Read(salt)
		for key, val := range salt {
			salt[key] = alphabet[val%byte(len(alphabet))]
		}
		e.salt = salt
	}
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
	return hash == hex.EncodeToString(pbkdf2.Key([]byte(rawData), e.salt, e.Iterations, e.KeyLen, e.HashFunc)), nil
}

// GetSalt Returns the salt if present, otherwise nil
func (e *Encoder) GetSalt() ([]byte, error) {
	return e.salt, nil
}
