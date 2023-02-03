package argon2

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"gopkg.in/encoder.v1/types"
	"strings"
)

const (
	DefaultMemory     uint32 = 64 * 1024
	DefaultIterations uint32 = 1
	DefaultThreads    uint8  = 4
)

type Encoder struct {
	// The amount of memory used by the algorithm (in kibibytes).
	Memory uint32
	// The number of iterations over the memory.
	Iterations uint32
	// The number of threads (or lanes) used by the algorithm.
	// Recommended value is between 1 and runtime.NumCPU().
	Threads uint8
	// Length of the random salt. 16 bytes is recommended for password hashing.
	SaltLen uint32
	// Length of the generated key. 16 bytes or more is recommended.
	KeyLen    uint32
	salt      []byte
	hashValue []byte
}

// Hash Generate and return a hash value in []byte format
func (e *Encoder) Hash(src string) ([]byte, error) {
	if e.hashValue != nil {
		return e.hashValue, nil
	}
	if e.salt == nil {
		salt, err := types.GenerateRandomSalt(int(e.SaltLen))
		if err != nil {
			return nil, err
		}
		e.salt = salt
	}
	e.hashValue = argon2.IDKey([]byte(src), e.salt, e.Iterations, e.Memory, e.Threads, e.KeyLen)
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
	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(e.salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(e.hashValue)
	// Return a string using the standard encoded hash format.
	encoded := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, e.Memory, e.Iterations, e.Threads, b64Salt, b64Hash)
	return encoded, nil
}

// Verify compares a encoded data with its possible plaintext equivalent
func (e *Encoder) Verify(hash, rawData string) (bool, error) {
	hashParts := strings.Split(hash, "$")
	_, err := fmt.Sscanf(hashParts[3], "m=%d,t=%d,p=%d", &e.Memory, &e.Iterations, &e.Threads)
	if err != nil {
		return false, err
	}
	salt, err := base64.RawStdEncoding.DecodeString(hashParts[4])
	if err != nil {
		return false, err
	}
	decodedHash, err := base64.RawStdEncoding.DecodeString(hashParts[5])
	if err != nil {
		return false, err
	}
	hashToCompare := argon2.IDKey([]byte(rawData), salt, e.Iterations, e.Memory, e.Threads, uint32(len(decodedHash)))
	return subtle.ConstantTimeCompare(decodedHash, hashToCompare) == 1, nil
}

// GetSalt Returns the salt if present, otherwise nil
func (e *Encoder) GetSalt() ([]byte, error) {
	return e.salt, nil
}
