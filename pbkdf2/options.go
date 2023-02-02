package pbkdf2

import (
	"gopkg.in/encoder.v1/types"
	"hash"
)

// WithSalt configure the salt value for Encoder, default automatically generate random strings
func WithSalt(salt []byte) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		p := e.(*Encoder)
		p.salt = salt
	})
}

// WithSaltLen configure the salt length for Encoder, default value is 16
func WithSaltLen(len int) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		p := e.(*Encoder)
		p.SaltLen = len
	})
}

// WithIterations configure the iterations for Encoder, default value is 10000
func WithIterations(iter int) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		p := e.(*Encoder)
		p.Iterations = iter
	})
}

// WithKeyLen configure the key length for Encoder, default value is 32
func WithKeyLen(len int) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		p := e.(*Encoder)
		p.KeyLen = len
	})
}

// WithHasFunc configure the ken hash func for Encoder, default value is sha256.New
func WithHasFunc(f func() hash.Hash) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		p := e.(*Encoder)
		p.HashFunc = f
	})
}
