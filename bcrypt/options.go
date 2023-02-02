package bcrypt

import "github.com/go-encoder/encoder/types"

// WithCost configure the cost for BcryptEncoder
func WithCost(cost int) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		b := e.(*Encoder)
		b.Cost = cost
	})
}
