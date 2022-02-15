// Code generated by a tool. DO NOT EDIT.

// Package macho provides a mockable wrapper for debug/macho.
package macho

import (
	macho "debug/macho"
	io "io"
)

var _ Interface = &Impl{}
var _ = macho.NewFatFile

type Interface interface {
	NewFatFile(r io.ReaderAt) (*macho.FatFile, error)
	NewFile(r io.ReaderAt) (*macho.File, error)
	Open(name string) (*macho.File, error)
	OpenFat(name string) (*macho.FatFile, error)
}

type Impl struct{}

func (*Impl) NewFatFile(r io.ReaderAt) (*macho.FatFile, error) {
	return macho.NewFatFile(r)
}
func (*Impl) NewFile(r io.ReaderAt) (*macho.File, error) {
	return macho.NewFile(r)
}
func (*Impl) Open(name string) (*macho.File, error) {
	return macho.Open(name)
}
func (*Impl) OpenFat(name string) (*macho.FatFile, error) {
	return macho.OpenFat(name)
}