// Code generated by a tool. DO NOT EDIT.

// Package mail provides a mockable wrapper for net/mail.
package mail

import (
	io "io"
	mail "net/mail"
	time "time"
)

var _ Interface = &Impl{}
var _ = mail.ParseAddress

type Interface interface {
	ParseAddress(address string) (*mail.Address, error)
	ParseAddressList(list string) ([]*mail.Address, error)
	ParseDate(date string) (time.Time, error)
	ReadMessage(r io.Reader) (msg *mail.Message, err error)
}

type Impl struct{}

func (*Impl) ParseAddress(address string) (*mail.Address, error) {
	return mail.ParseAddress(address)
}
func (*Impl) ParseAddressList(list string) ([]*mail.Address, error) {
	return mail.ParseAddressList(list)
}
func (*Impl) ParseDate(date string) (time.Time, error) {
	return mail.ParseDate(date)
}
func (*Impl) ReadMessage(r io.Reader) (msg *mail.Message, err error) {
	return mail.ReadMessage(r)
}
