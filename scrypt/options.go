package scrypt

import "github.com/go-encoder/encoder/types"

// WithSaltLen configure the salt length for Encoder, default value is 16
func WithSaltLen(len int) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		b := e.(*Encoder)
		b.SaltLen = len
	})
}

// WithSalt configure the salt value for Encoder, default automatically generate random strings
func WithSalt(salt []byte) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		b := e.(*Encoder)
		b.salt = salt
	})
}

// WithN configure the N value for Encoder, default value is 32768
func WithN(N int) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		b := e.(*Encoder)
		b.N = N
	})
}

// WithR configure the R value for Encoder, default value is 8
func WithR(R int) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		b := e.(*Encoder)
		b.R = R
	})
}

// WithP configure the P value for Encoder, default value is 8
func WithP(P int) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		b := e.(*Encoder)
		b.P = P
	})
}
