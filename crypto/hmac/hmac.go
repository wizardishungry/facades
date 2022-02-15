// Code generated by a tool. DO NOT EDIT.

// Package hmac provides a mockable wrapper for crypto/hmac.
package hmac

import (
	hmac "crypto/hmac"
	hash "hash"
)

var _ Interface = &Impl{}
var _ = hmac.Equal

type Interface interface {
	Equal(mac1 []byte, mac2 []byte) bool
	New(h func() hash.Hash, key []byte) hash.Hash
}

type Impl struct{}

func (*Impl) Equal(mac1 []byte, mac2 []byte) bool {
	return hmac.Equal(mac1, mac2)
}
func (*Impl) New(h func() hash.Hash, key []byte) hash.Hash {
	return hmac.New(h, key)
}
