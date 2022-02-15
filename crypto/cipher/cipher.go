// Code generated by a tool. DO NOT EDIT.

// Package cipher provides a mockable wrapper for crypto/cipher.
package cipher

import (
	cipher "crypto/cipher"
)

var _ Interface = &Impl{}
var _ = cipher.NewCBCDecrypter

type Interface interface {
	NewCBCDecrypter(b cipher.Block, iv []byte) cipher.BlockMode
	NewCBCEncrypter(b cipher.Block, iv []byte) cipher.BlockMode
	NewCFBDecrypter(block cipher.Block, iv []byte) cipher.Stream
	NewCFBEncrypter(block cipher.Block, iv []byte) cipher.Stream
	NewCTR(block cipher.Block, iv []byte) cipher.Stream
	NewGCM(cipher cipher.Block) (cipher.AEAD, error)
	NewGCMWithNonceSize(cipher cipher.Block, size int) (cipher.AEAD, error)
	NewGCMWithTagSize(cipher cipher.Block, tagSize int) (cipher.AEAD, error)
	NewOFB(b cipher.Block, iv []byte) cipher.Stream
}

type Impl struct{}

func (*Impl) NewCBCDecrypter(b cipher.Block, iv []byte) cipher.BlockMode {
	return cipher.NewCBCDecrypter(b, iv)
}
func (*Impl) NewCBCEncrypter(b cipher.Block, iv []byte) cipher.BlockMode {
	return cipher.NewCBCEncrypter(b, iv)
}
func (*Impl) NewCFBDecrypter(block cipher.Block, iv []byte) cipher.Stream {
	return cipher.NewCFBDecrypter(block, iv)
}
func (*Impl) NewCFBEncrypter(block cipher.Block, iv []byte) cipher.Stream {
	return cipher.NewCFBEncrypter(block, iv)
}
func (*Impl) NewCTR(block cipher.Block, iv []byte) cipher.Stream {
	return cipher.NewCTR(block, iv)
}
func (*Impl) NewGCM(cipher_v cipher.Block) (cipher.AEAD, error) {
	return cipher.NewGCM(cipher_v)
}
func (*Impl) NewGCMWithNonceSize(cipher_v cipher.Block, size_v int) (cipher.AEAD, error) {
	return cipher.NewGCMWithNonceSize(cipher_v, size_v)
}
func (*Impl) NewGCMWithTagSize(cipher_v cipher.Block, tagSize_v int) (cipher.AEAD, error) {
	return cipher.NewGCMWithTagSize(cipher_v, tagSize_v)
}
func (*Impl) NewOFB(b cipher.Block, iv []byte) cipher.Stream {
	return cipher.NewOFB(b, iv)
}
