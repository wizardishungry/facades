// Code generated by a tool. DO NOT EDIT.

// Package quick provides a mockable wrapper for testing/quick.
package quick

import (
	rand "math/rand"
	reflect "reflect"
	quick "testing/quick"
)

var _ Interface = &Impl{}
var _ = quick.Check

type Interface interface {
	Check(f any, config *quick.Config) error
	CheckEqual(f any, g any, config *quick.Config) error
	Value(t reflect.Type, rand *rand.Rand) (value reflect.Value, ok bool)
}

type Impl struct{}

func (*Impl) Check(f any, config *quick.Config) error {
	return quick.Check(f, config)
}
func (*Impl) CheckEqual(f any, g any, config *quick.Config) error {
	return quick.CheckEqual(f, g, config)
}
func (*Impl) Value(t reflect.Type, rand *rand.Rand) (value reflect.Value, ok bool) {
	return quick.Value(t, rand)
}
