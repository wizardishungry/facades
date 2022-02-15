// Code generated by a tool. DO NOT EDIT.

// Package ioutil provides a mockable wrapper for io/ioutil.
package ioutil

import (
	io "io"
	fs "io/fs"
	ioutil "io/ioutil"
	os "os"
)

var _ Interface = &Impl{}
var _ = ioutil.NopCloser

type Interface interface {
	NopCloser(r io.Reader) io.ReadCloser
	ReadAll(r io.Reader) ([]byte, error)
	ReadDir(dirname string) ([]fs.FileInfo, error)
	ReadFile(filename string) ([]byte, error)
	TempDir(dir string, pattern string) (name string, err error)
	TempFile(dir string, pattern string) (f *os.File, err error)
	WriteFile(filename string, data []byte, perm fs.FileMode) error
}

type Impl struct{}

func (*Impl) NopCloser(r io.Reader) io.ReadCloser {
	return ioutil.NopCloser(r)
}
func (*Impl) ReadAll(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}
func (*Impl) ReadDir(dirname string) ([]fs.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}
func (*Impl) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
func (*Impl) TempDir(dir string, pattern string) (name string, err error) {
	return ioutil.TempDir(dir, pattern)
}
func (*Impl) TempFile(dir string, pattern string) (f *os.File, err error) {
	return ioutil.TempFile(dir, pattern)
}
func (*Impl) WriteFile(filename string, data []byte, perm fs.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}
