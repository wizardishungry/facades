// Code generated by a tool. DO NOT EDIT.

// Package adler32 provides a mockable wrapper for hash/adler32.
package adler32

import (
	hash "hash"
	adler32 "hash/adler32"
)

var _ Interface = &Impl{}
var _ = adler32.Checksum

type Interface interface {
	Checksum(data []byte) uint32
	New() hash.Hash32
}

type Impl struct{}

func (*Impl) Checksum(data []byte) uint32 {
	return adler32.Checksum(data)
}
func (*Impl) New() hash.Hash32 {
	return adler32.New()
}
