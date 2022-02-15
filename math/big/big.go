// Code generated by a tool. DO NOT EDIT.

// Package big provides a mockable wrapper for math/big.
package big

import (
	big "math/big"
)

var _ Interface = &Impl{}
var _ = big.Jacobi

type Interface interface {
	Jacobi(x *big.Int, y *big.Int) int
	NewFloat(x float64) *big.Float
	NewInt(x int64) *big.Int
	NewRat(a int64, b int64) *big.Rat
	ParseFloat(s string, base int, prec uint, mode big.RoundingMode) (f *big.Float, b int, err error)
}

type Impl struct{}

func (*Impl) Jacobi(x *big.Int, y *big.Int) int {
	return big.Jacobi(x, y)
}
func (*Impl) NewFloat(x float64) *big.Float {
	return big.NewFloat(x)
}
func (*Impl) NewInt(x int64) *big.Int {
	return big.NewInt(x)
}
func (*Impl) NewRat(a int64, b int64) *big.Rat {
	return big.NewRat(a, b)
}
func (*Impl) ParseFloat(s string, base int, prec uint, mode big.RoundingMode) (f *big.Float, b int, err error) {
	return big.ParseFloat(s, base, prec, mode)
}
