// Code generated by a tool. DO NOT EDIT.

// Package multipart provides a mockable wrapper for mime/multipart.
package multipart

import (
	io "io"
	multipart "mime/multipart"
)

var _ Interface = &Impl{}
var _ = multipart.NewReader

type Interface interface {
	NewReader(r io.Reader, boundary string) *multipart.Reader
	NewWriter(w io.Writer) *multipart.Writer
}

type Impl struct{}

func (*Impl) NewReader(r io.Reader, boundary string) *multipart.Reader {
	return multipart.NewReader(r, boundary)
}
func (*Impl) NewWriter(w io.Writer) *multipart.Writer {
	return multipart.NewWriter(w)
}
