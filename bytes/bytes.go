// Code generated by a tool. DO NOT EDIT.

// Package bytes provides a mockable wrapper for bytes.
package bytes

import (
	bytes "bytes"
	unicode "unicode"
)

var _ Interface = &Impl{}
var _ = bytes.Compare

type Interface interface {
	Compare(a []byte, b []byte) int
	Contains(b []byte, subslice []byte) bool
	ContainsAny(b []byte, chars string) bool
	ContainsRune(b []byte, r rune) bool
	Count(s []byte, sep []byte) int
	Cut(s []byte, sep []byte) (before []byte, after []byte, found bool)
	Equal(a []byte, b []byte) bool
	EqualFold(s []byte, t []byte) bool
	Fields(s []byte) [][]byte
	FieldsFunc(s []byte, f func(rune) bool) [][]byte
	HasPrefix(s []byte, prefix []byte) bool
	HasSuffix(s []byte, suffix []byte) bool
	Index(s []byte, sep []byte) int
	IndexAny(s []byte, chars string) int
	IndexByte(b []byte, c byte) int
	IndexFunc(s []byte, f func(r rune) bool) int
	IndexRune(s []byte, r rune) int
	Join(s [][]byte, sep []byte) []byte
	LastIndex(s []byte, sep []byte) int
	LastIndexAny(s []byte, chars string) int
	LastIndexByte(s []byte, c byte) int
	LastIndexFunc(s []byte, f func(r rune) bool) int
	Map(mapping func(r rune) rune, s []byte) []byte
	NewBuffer(buf []byte) *bytes.Buffer
	NewBufferString(s string) *bytes.Buffer
	NewReader(b []byte) *bytes.Reader
	Repeat(b []byte, count int) []byte
	Replace(s []byte, old []byte, new []byte, n int) []byte
	ReplaceAll(s []byte, old []byte, new []byte) []byte
	Runes(s []byte) []rune
	Split(s []byte, sep []byte) [][]byte
	SplitAfter(s []byte, sep []byte) [][]byte
	SplitAfterN(s []byte, sep []byte, n int) [][]byte
	SplitN(s []byte, sep []byte, n int) [][]byte
	Title(s []byte) []byte
	ToLower(s []byte) []byte
	ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte
	ToTitle(s []byte) []byte
	ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte
	ToUpper(s []byte) []byte
	ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte
	ToValidUTF8(s []byte, replacement []byte) []byte
	Trim(s []byte, cutset string) []byte
	TrimFunc(s []byte, f func(r rune) bool) []byte
	TrimLeft(s []byte, cutset string) []byte
	TrimLeftFunc(s []byte, f func(r rune) bool) []byte
	TrimPrefix(s []byte, prefix []byte) []byte
	TrimRight(s []byte, cutset string) []byte
	TrimRightFunc(s []byte, f func(r rune) bool) []byte
	TrimSpace(s []byte) []byte
	TrimSuffix(s []byte, suffix []byte) []byte
}

type Impl struct{}

func (*Impl) Compare(a []byte, b []byte) int {
	return bytes.Compare(a, b)
}
func (*Impl) Contains(b []byte, subslice []byte) bool {
	return bytes.Contains(b, subslice)
}
func (*Impl) ContainsAny(b []byte, chars string) bool {
	return bytes.ContainsAny(b, chars)
}
func (*Impl) ContainsRune(b []byte, r rune) bool {
	return bytes.ContainsRune(b, r)
}
func (*Impl) Count(s []byte, sep []byte) int {
	return bytes.Count(s, sep)
}
func (*Impl) Cut(s []byte, sep []byte) (before []byte, after []byte, found bool) {
	return bytes.Cut(s, sep)
}
func (*Impl) Equal(a []byte, b []byte) bool {
	return bytes.Equal(a, b)
}
func (*Impl) EqualFold(s []byte, t []byte) bool {
	return bytes.EqualFold(s, t)
}
func (*Impl) Fields(s []byte) [][]byte {
	return bytes.Fields(s)
}
func (*Impl) FieldsFunc(s []byte, f func(rune) bool) [][]byte {
	return bytes.FieldsFunc(s, f)
}
func (*Impl) HasPrefix(s []byte, prefix []byte) bool {
	return bytes.HasPrefix(s, prefix)
}
func (*Impl) HasSuffix(s []byte, suffix []byte) bool {
	return bytes.HasSuffix(s, suffix)
}
func (*Impl) Index(s []byte, sep []byte) int {
	return bytes.Index(s, sep)
}
func (*Impl) IndexAny(s []byte, chars string) int {
	return bytes.IndexAny(s, chars)
}
func (*Impl) IndexByte(b []byte, c byte) int {
	return bytes.IndexByte(b, c)
}
func (*Impl) IndexFunc(s []byte, f func(r rune) bool) int {
	return bytes.IndexFunc(s, f)
}
func (*Impl) IndexRune(s []byte, r rune) int {
	return bytes.IndexRune(s, r)
}
func (*Impl) Join(s [][]byte, sep []byte) []byte {
	return bytes.Join(s, sep)
}
func (*Impl) LastIndex(s []byte, sep []byte) int {
	return bytes.LastIndex(s, sep)
}
func (*Impl) LastIndexAny(s []byte, chars string) int {
	return bytes.LastIndexAny(s, chars)
}
func (*Impl) LastIndexByte(s []byte, c byte) int {
	return bytes.LastIndexByte(s, c)
}
func (*Impl) LastIndexFunc(s []byte, f func(r rune) bool) int {
	return bytes.LastIndexFunc(s, f)
}
func (*Impl) Map(mapping func(r rune) rune, s []byte) []byte {
	return bytes.Map(mapping, s)
}
func (*Impl) NewBuffer(buf []byte) *bytes.Buffer {
	return bytes.NewBuffer(buf)
}
func (*Impl) NewBufferString(s string) *bytes.Buffer {
	return bytes.NewBufferString(s)
}
func (*Impl) NewReader(b []byte) *bytes.Reader {
	return bytes.NewReader(b)
}
func (*Impl) Repeat(b []byte, count int) []byte {
	return bytes.Repeat(b, count)
}
func (*Impl) Replace(s []byte, old []byte, new []byte, n int) []byte {
	return bytes.Replace(s, old, new, n)
}
func (*Impl) ReplaceAll(s []byte, old []byte, new []byte) []byte {
	return bytes.ReplaceAll(s, old, new)
}
func (*Impl) Runes(s []byte) []rune {
	return bytes.Runes(s)
}
func (*Impl) Split(s []byte, sep []byte) [][]byte {
	return bytes.Split(s, sep)
}
func (*Impl) SplitAfter(s []byte, sep []byte) [][]byte {
	return bytes.SplitAfter(s, sep)
}
func (*Impl) SplitAfterN(s []byte, sep []byte, n int) [][]byte {
	return bytes.SplitAfterN(s, sep, n)
}
func (*Impl) SplitN(s []byte, sep []byte, n int) [][]byte {
	return bytes.SplitN(s, sep, n)
}
func (*Impl) Title(s []byte) []byte {
	return bytes.Title(s)
}
func (*Impl) ToLower(s []byte) []byte {
	return bytes.ToLower(s)
}
func (*Impl) ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte {
	return bytes.ToLowerSpecial(c, s)
}
func (*Impl) ToTitle(s []byte) []byte {
	return bytes.ToTitle(s)
}
func (*Impl) ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte {
	return bytes.ToTitleSpecial(c, s)
}
func (*Impl) ToUpper(s []byte) []byte {
	return bytes.ToUpper(s)
}
func (*Impl) ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte {
	return bytes.ToUpperSpecial(c, s)
}
func (*Impl) ToValidUTF8(s []byte, replacement []byte) []byte {
	return bytes.ToValidUTF8(s, replacement)
}
func (*Impl) Trim(s []byte, cutset string) []byte {
	return bytes.Trim(s, cutset)
}
func (*Impl) TrimFunc(s []byte, f func(r rune) bool) []byte {
	return bytes.TrimFunc(s, f)
}
func (*Impl) TrimLeft(s []byte, cutset string) []byte {
	return bytes.TrimLeft(s, cutset)
}
func (*Impl) TrimLeftFunc(s []byte, f func(r rune) bool) []byte {
	return bytes.TrimLeftFunc(s, f)
}
func (*Impl) TrimPrefix(s []byte, prefix []byte) []byte {
	return bytes.TrimPrefix(s, prefix)
}
func (*Impl) TrimRight(s []byte, cutset string) []byte {
	return bytes.TrimRight(s, cutset)
}
func (*Impl) TrimRightFunc(s []byte, f func(r rune) bool) []byte {
	return bytes.TrimRightFunc(s, f)
}
func (*Impl) TrimSpace(s []byte) []byte {
	return bytes.TrimSpace(s)
}
func (*Impl) TrimSuffix(s []byte, suffix []byte) []byte {
	return bytes.TrimSuffix(s, suffix)
}
