package bcrypt

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	MinCost     int = 4  // the minimum allowable cost as passed in to Encode string
	MaxCost     int = 31 // the maximum allowable cost as passed in to Encode string
	DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into Encode string
)

type Encoder struct {
	Cost      int
	hashValue []byte
}

// Hash Generate and return a hash value in []byte format
func (e *Encoder) Hash(src string) ([]byte, error) {
	if e.hashValue != nil {
		return e.hashValue, nil
	}
	var err error
	e.hashValue, err = bcrypt.GenerateFromPassword([]byte(src), e.Cost)
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
	return string(e.hashValue), nil
}

// Verify compares a encoded data with its possible plaintext equivalent
func (e *Encoder) Verify(hash, rawData string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(rawData))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false, nil
	}
	return err == nil, err
}

// GetSalt Returns the salt if present, otherwise nil
func (e *Encoder) GetSalt() ([]byte, error) {
	return nil, nil
}
