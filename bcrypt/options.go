package bcrypt

import "gopkg.in/encoder.v1/types"

// WithCost configure the cost for BcryptEncoder
func WithCost(cost int) types.Option {
	return types.OptionFunc(func(e types.Encoder) {
		b := e.(*Encoder)
		b.Cost = cost
	})
}
