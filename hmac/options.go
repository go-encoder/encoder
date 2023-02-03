package hmac

import (
	"gopkg.in/encoder.v1/types"
	"hash"
)

// WithKey configure the hash key for Encoder, default value is ""
func WithKey(key string) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		p := e.(*Encoder)
		p.key = key
	})
}

// WithHasFunc configure the ken hash func for Encoder, default value is sha256.New
func WithHasFunc(f func() hash.Hash) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		p := e.(*Encoder)
		p.HashFunc = f
	})
}
