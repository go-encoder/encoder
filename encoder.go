package encoder

import (
	"github.com/go-encoder/encoder/argon2"
	"github.com/go-encoder/encoder/bcrypt"
	"github.com/go-encoder/encoder/hkdf"
	"github.com/go-encoder/encoder/pbkdf2"
	"github.com/go-encoder/encoder/scrypt"
	"github.com/go-encoder/encoder/types"
)

// New Returns an encoder instance of the specified type
func New(t types.EncoderType, opts ...types.Option) types.Encoder {
	var e types.Encoder
	switch t {
	case types.Bcrypt:
		e = &bcrypt.Encoder{Cost: bcrypt.DefaultCost}
	case types.Pbkdf2:
		e = &pbkdf2.Encoder{
			SaltLen:    types.SaltLen,
			Iterations: pbkdf2.DefaultIterations,
			KeyLen:     types.DefaultKeyLen,
			HashFunc:   types.DefaultHashFunc,
		}
	case types.Argon2:
		e = &argon2.Encoder{
			Memory:  argon2.DefaultMemory,
			Time:    argon2.DefaultTime,
			KeyLen:  uint32(types.DefaultKeyLen),
			SaltLen: uint32(types.SaltLen),
			Threads: argon2.DefaultThreads,
		}
	case types.Scrypt:
		e = &scrypt.Encoder{
			SaltLen: types.SaltLen * 2,
			N:       1 << 15,
			R:       8,
			P:       1,
		}
	case types.Hkdf:
		e = &hkdf.Encoder{
			SaltLen:  types.SaltLen,
			HashFunc: types.DefaultHashFunc,
			Info:     "",
			HashLen:  types.DefaultKeyLen,
		}
	}
	for _, opt := range opts {
		opt.Apply(e)
	}
	return e
}

// NewBcryptEncoder Returns Bcrypt encoder instance
func NewBcryptEncoder(opts ...types.Option) types.Encoder {
	return New(types.Bcrypt, opts...)
}

// NewPbkdf2Encoder Returns Pbkdf2 encoder instance
func NewPbkdf2Encoder(opts ...types.Option) types.Encoder {
	return New(types.Pbkdf2, opts...)
}

// NewArgon2Encoder Returns Argon2 encoder instance
func NewArgon2Encoder(opts ...types.Option) types.Encoder {
	return New(types.Argon2, opts...)
}

// NewScryptEncoder Returns Scrypt encoder instance
func NewScryptEncoder(opts ...types.Option) types.Encoder {
	return New(types.Scrypt, opts...)
}