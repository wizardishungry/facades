// Code generated by a tool. DO NOT EDIT.

// Package cmplx provides a mockable wrapper for math/cmplx.
package cmplx

import (
	cmplx "math/cmplx"
)

var _ Interface = &Impl{}
var _ = cmplx.Abs

type Interface interface {
	Abs(x complex128) float64
	Acos(x complex128) complex128
	Acosh(x complex128) complex128
	Asin(x complex128) complex128
	Asinh(x complex128) complex128
	Atan(x complex128) complex128
	Atanh(x complex128) complex128
	Conj(x complex128) complex128
	Cos(x complex128) complex128
	Cosh(x complex128) complex128
	Cot(x complex128) complex128
	Exp(x complex128) complex128
	Inf() complex128
	IsInf(x complex128) bool
	IsNaN(x complex128) bool
	Log(x complex128) complex128
	Log10(x complex128) complex128
	NaN() complex128
	Phase(x complex128) float64
	Polar(x complex128) (r float64, θ float64)
	Pow(x complex128, y complex128) complex128
	Rect(r float64, θ float64) complex128
	Sin(x complex128) complex128
	Sinh(x complex128) complex128
	Sqrt(x complex128) complex128
	Tan(x complex128) complex128
	Tanh(x complex128) complex128
}

type Impl struct{}

func (*Impl) Abs(x complex128) float64 {
	return cmplx.Abs(x)
}
func (*Impl) Acos(x complex128) complex128 {
	return cmplx.Acos(x)
}
func (*Impl) Acosh(x complex128) complex128 {
	return cmplx.Acosh(x)
}
func (*Impl) Asin(x complex128) complex128 {
	return cmplx.Asin(x)
}
func (*Impl) Asinh(x complex128) complex128 {
	return cmplx.Asinh(x)
}
func (*Impl) Atan(x complex128) complex128 {
	return cmplx.Atan(x)
}
func (*Impl) Atanh(x complex128) complex128 {
	return cmplx.Atanh(x)
}
func (*Impl) Conj(x complex128) complex128 {
	return cmplx.Conj(x)
}
func (*Impl) Cos(x complex128) complex128 {
	return cmplx.Cos(x)
}
func (*Impl) Cosh(x complex128) complex128 {
	return cmplx.Cosh(x)
}
func (*Impl) Cot(x complex128) complex128 {
	return cmplx.Cot(x)
}
func (*Impl) Exp(x complex128) complex128 {
	return cmplx.Exp(x)
}
func (*Impl) Inf() complex128 {
	return cmplx.Inf()
}
func (*Impl) IsInf(x complex128) bool {
	return cmplx.IsInf(x)
}
func (*Impl) IsNaN(x complex128) bool {
	return cmplx.IsNaN(x)
}
func (*Impl) Log(x complex128) complex128 {
	return cmplx.Log(x)
}
func (*Impl) Log10(x complex128) complex128 {
	return cmplx.Log10(x)
}
func (*Impl) NaN() complex128 {
	return cmplx.NaN()
}
func (*Impl) Phase(x complex128) float64 {
	return cmplx.Phase(x)
}
func (*Impl) Polar(x complex128) (r float64, θ float64) {
	return cmplx.Polar(x)
}
func (*Impl) Pow(x complex128, y complex128) complex128 {
	return cmplx.Pow(x, y)
}
func (*Impl) Rect(r float64, θ float64) complex128 {
	return cmplx.Rect(r, θ)
}
func (*Impl) Sin(x complex128) complex128 {
	return cmplx.Sin(x)
}
func (*Impl) Sinh(x complex128) complex128 {
	return cmplx.Sinh(x)
}
func (*Impl) Sqrt(x complex128) complex128 {
	return cmplx.Sqrt(x)
}
func (*Impl) Tan(x complex128) complex128 {
	return cmplx.Tan(x)
}
func (*Impl) Tanh(x complex128) complex128 {
	return cmplx.Tanh(x)
}
