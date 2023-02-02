package types

import (
	"crypto/rand"
	"crypto/sha256"
)

type EncoderType string

const (
	Bcrypt        EncoderType = "bcrypt" // use bcrypt hash
	Scrypt        EncoderType = "scrypt" // use scrypt hash
	Pbkdf2        EncoderType = "pbkdf2" // use pbkdf2 hash
	Argon2        EncoderType = "argon2" // use argon2 hash
	Hkdf          EncoderType = "hkdf"   // use hkdf hash
	SaltLen       int         = 16       // default size of salt value
	DefaultKeyLen             = 32
)

var DefaultHashFunc = sha256.New

type Encoder interface {
	// Encode returns the hash value of the given data
	Encode(src string) (string, error)
	// Verify compares a encoded data with its possible plaintext equivalent
	Verify(hash string, rawData string) (bool, error)
	// GetSalt Returns the salt if exists, otherwise nil
	GetSalt() ([]byte, error)
}

// OptionFunc wraps a func so it satisfies the Option interface.
type OptionFunc func(e Encoder)

// Apply config options for instance
func (f OptionFunc) Apply(e Encoder) {
	f(e)
}

// An Option configures a Encoder
type Option interface {
	Apply(Encoder)
}

// GenerateRandomSalt generates random salt.
func GenerateRandomSalt(len int) ([]byte, error) {
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
