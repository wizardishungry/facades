// Code generated by a tool. DO NOT EDIT.

// Package fcgi provides a mockable wrapper for net/http/fcgi.
package fcgi

import (
	net "net"
	http "net/http"
	fcgi "net/http/fcgi"
)

var _ Interface = &Impl{}
var _ = fcgi.ProcessEnv

type Interface interface {
	ProcessEnv(r *http.Request) map[string]string
	Serve(l net.Listener, handler http.Handler) error
}

type Impl struct{}

func (*Impl) ProcessEnv(r *http.Request) map[string]string {
	return fcgi.ProcessEnv(r)
}
func (*Impl) Serve(l net.Listener, handler http.Handler) error {
	return fcgi.Serve(l, handler)
}
