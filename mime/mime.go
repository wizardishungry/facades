// Code generated by a tool. DO NOT EDIT.

// Package mime provides a mockable wrapper for mime.
package mime

import (
	mime "mime"
)

var _ Interface = &Impl{}
var _ = mime.AddExtensionType

type Interface interface {
	AddExtensionType(ext string, typ string) error
	ExtensionsByType(typ string) ([]string, error)
	FormatMediaType(t string, param map[string]string) string
	ParseMediaType(v string) (mediatype string, params map[string]string, err error)
	TypeByExtension(ext string) string
}

type Impl struct{}

func (*Impl) AddExtensionType(ext string, typ string) error {
	return mime.AddExtensionType(ext, typ)
}
func (*Impl) ExtensionsByType(typ string) ([]string, error) {
	return mime.ExtensionsByType(typ)
}
func (*Impl) FormatMediaType(t string, param map[string]string) string {
	return mime.FormatMediaType(t, param)
}
func (*Impl) ParseMediaType(v string) (mediatype string, params map[string]string, err error) {
	return mime.ParseMediaType(v)
}
func (*Impl) TypeByExtension(ext string) string {
	return mime.TypeByExtension(ext)
}
