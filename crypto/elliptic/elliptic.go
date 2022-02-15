// Code generated by a tool. DO NOT EDIT.

// Package elliptic provides a mockable wrapper for crypto/elliptic.
package elliptic

import (
	elliptic "crypto/elliptic"
	io "io"
	big "math/big"
)

var _ Interface = &Impl{}
var _ = elliptic.GenerateKey

type Interface interface {
	GenerateKey(curve elliptic.Curve, rand io.Reader) (priv []byte, x *big.Int, y *big.Int, err error)
	Marshal(curve elliptic.Curve, x *big.Int, y *big.Int) []byte
	MarshalCompressed(curve elliptic.Curve, x *big.Int, y *big.Int) []byte
	P224() elliptic.Curve
	P256() elliptic.Curve
	P384() elliptic.Curve
	P521() elliptic.Curve
	Unmarshal(curve elliptic.Curve, data []byte) (x *big.Int, y *big.Int)
	UnmarshalCompressed(curve elliptic.Curve, data []byte) (x *big.Int, y *big.Int)
}

type Impl struct{}

func (*Impl) GenerateKey(curve elliptic.Curve, rand io.Reader) (priv []byte, x *big.Int, y *big.Int, err error) {
	return elliptic.GenerateKey(curve, rand)
}
func (*Impl) Marshal(curve elliptic.Curve, x *big.Int, y *big.Int) []byte {
	return elliptic.Marshal(curve, x, y)
}
func (*Impl) MarshalCompressed(curve elliptic.Curve, x *big.Int, y *big.Int) []byte {
	return elliptic.MarshalCompressed(curve, x, y)
}
func (*Impl) P224() elliptic.Curve {
	return elliptic.P224()
}
func (*Impl) P256() elliptic.Curve {
	return elliptic.P256()
}
func (*Impl) P384() elliptic.Curve {
	return elliptic.P384()
}
func (*Impl) P521() elliptic.Curve {
	return elliptic.P521()
}
func (*Impl) Unmarshal(curve elliptic.Curve, data []byte) (x *big.Int, y *big.Int) {
	return elliptic.Unmarshal(curve, data)
}
func (*Impl) UnmarshalCompressed(curve elliptic.Curve, data []byte) (x *big.Int, y *big.Int) {
	return elliptic.UnmarshalCompressed(curve, data)
}
