package argon2

import "gopkg.in/encoder.v1/types"

// WithMemory configure the memory for Encoder, default value is 64 * 1024
func WithMemory(memory uint32) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		b := e.(*Encoder)
		b.Memory = memory
	})
}

// WithIterations configure the iterations for Encoder, default value is 1
func WithIterations(iterations uint32) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		b := e.(*Encoder)
		b.Iterations = iterations
	})
}

// WithThreads configure the Threads for Encoder, default value is 4
func WithThreads(threads uint8) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		b := e.(*Encoder)
		b.Threads = threads
	})
}

// WithSaltLen configure the salt length for Encoder, default value is 16
func WithSaltLen(len uint32) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		b := e.(*Encoder)
		b.SaltLen = len
	})
}

// WithKeyLen configure the key length for Encoder, default value is 32
func WithKeyLen(len uint32) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		b := e.(*Encoder)
		b.KeyLen = len
	})
}

// WithSalt configure the salt value for Encoder, default automatically generate random strings
func WithSalt(salt []byte) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		b := e.(*Encoder)
		b.salt = salt
	})
}
