// Code generated by a tool. DO NOT EDIT.

// Package smtp provides a mockable wrapper for net/smtp.
package smtp

import (
	net "net"
	smtp "net/smtp"
)

var _ Interface = &Impl{}
var _ = smtp.CRAMMD5Auth

type Interface interface {
	CRAMMD5Auth(username string, secret string) smtp.Auth
	Dial(addr string) (*smtp.Client, error)
	NewClient(conn net.Conn, host string) (*smtp.Client, error)
	PlainAuth(identity string, username string, password string, host string) smtp.Auth
	SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error
}

type Impl struct{}

func (*Impl) CRAMMD5Auth(username string, secret string) smtp.Auth {
	return smtp.CRAMMD5Auth(username, secret)
}
func (*Impl) Dial(addr string) (*smtp.Client, error) {
	return smtp.Dial(addr)
}
func (*Impl) NewClient(conn net.Conn, host string) (*smtp.Client, error) {
	return smtp.NewClient(conn, host)
}
func (*Impl) PlainAuth(identity string, username string, password string, host string) smtp.Auth {
	return smtp.PlainAuth(identity, username, password, host)
}
func (*Impl) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return smtp.SendMail(addr, a, from, to, msg)
}
