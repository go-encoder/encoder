package hkdf

import (
	"github.com/go-encoder/encoder/types"
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

// WithHashLen configure the hash length for Encoder, default value is 32
func WithHashLen(len int) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		p := e.(*Encoder)
		p.HashLen = len
	})
}

// WithHasFunc configure the ken hash func for Encoder, default value is sha256.New
func WithHasFunc(f func() hash.Hash) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		p := e.(*Encoder)
		p.HashFunc = f
	})
}

// WithInfo configure the info value for Encoder, default value is ""
func WithInfo(info string) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		p := e.(*Encoder)
		p.Info = info
	})
}
