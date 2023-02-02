package argon2

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"github.com/go-encoder/encoder/types"
	"golang.org/x/crypto/argon2"
	"strings"
)

const (
	DefaultMemory  uint32 = 64 * 1024
	DefaultTime    uint32 = 1
	DefaultThreads uint8  = 4
)

type Encoder struct {
	Memory  uint32
	Time    uint32
	Threads uint8
	SaltLen uint32
	KeyLen  uint32
	salt    []byte
}

// Encode returns the hash value of the given data
func (e *Encoder) Encode(src string) (string, error) {
	if e.salt == nil {
		salt, err := types.GenerateRandomSalt(int(e.SaltLen))
		if err != nil {
			return "", err
		}
		e.salt = salt
	}
	hash := argon2.IDKey([]byte(src), e.salt, e.Time, e.Memory, e.Threads, e.KeyLen)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(e.salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	// Return a string using the standard encoded hash format.
	encoded := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, e.Memory, e.Time, e.Threads, b64Salt, b64Hash)
	return encoded, nil
}

// Verify compares a encoded data with its possible plaintext equivalent
func (e *Encoder) Verify(hash, rawData string) (bool, error) {
	hashParts := strings.Split(hash, "$")
	_, err := fmt.Sscanf(hashParts[3], "m=%d,t=%d,p=%d", &e.Memory, &e.Time, &e.Threads)
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
	hashToCompare := argon2.IDKey([]byte(rawData), salt, e.Time, e.Memory, e.Threads, uint32(len(decodedHash)))
	return subtle.ConstantTimeCompare(decodedHash, hashToCompare) == 1, nil
}

// GetSalt Returns the salt if present, otherwise nil
func (e *Encoder) GetSalt() ([]byte, error) {
	return e.salt, nil
}
