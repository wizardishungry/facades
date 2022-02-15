// Code generated by a tool. DO NOT EDIT.

// Package plan9obj provides a mockable wrapper for debug/plan9obj.
package plan9obj

import (
	plan9obj "debug/plan9obj"
	io "io"
)

var _ Interface = &Impl{}
var _ = plan9obj.NewFile

type Interface interface {
	NewFile(r io.ReaderAt) (*plan9obj.File, error)
	Open(name string) (*plan9obj.File, error)
}

type Impl struct{}

func (*Impl) NewFile(r io.ReaderAt) (*plan9obj.File, error) {
	return plan9obj.NewFile(r)
}
func (*Impl) Open(name string) (*plan9obj.File, error) {
	return plan9obj.Open(name)
}