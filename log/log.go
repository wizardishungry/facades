// Code generated by a tool. DO NOT EDIT.

// Package log provides a mockable wrapper for log.
package log

import (
	io "io"
	log "log"
)

var _ Interface = &Impl{}
var _ = log.Default

type Interface interface {
	Default() *log.Logger
	Fatal(v ...any)
	Fatalf(format string, v ...any)
	Fatalln(v ...any)
	Flags() int
	New(out io.Writer, prefix string, flag int) *log.Logger
	Output(calldepth int, s string) error
	Panic(v ...any)
	Panicf(format string, v ...any)
	Panicln(v ...any)
	Prefix() string
	Print(v ...any)
	Printf(format string, v ...any)
	Println(v ...any)
	SetFlags(flag int)
	SetOutput(w io.Writer)
	SetPrefix(prefix string)
	Writer() io.Writer
}

type Impl struct{}

func (*Impl) Default() *log.Logger {
	return log.Default()
}
func (*Impl) Fatal(v ...any) {
	log.Fatal(v...)
}
func (*Impl) Fatalf(format string, v ...any) {
	log.Fatalf(format, v...)
}
func (*Impl) Fatalln(v ...any) {
	log.Fatalln(v...)
}
func (*Impl) Flags() int {
	return log.Flags()
}
func (*Impl) New(out io.Writer, prefix string, flag int) *log.Logger {
	return log.New(out, prefix, flag)
}
func (*Impl) Output(calldepth int, s string) error {
	return log.Output(calldepth, s)
}
func (*Impl) Panic(v ...any) {
	log.Panic(v...)
}
func (*Impl) Panicf(format string, v ...any) {
	log.Panicf(format, v...)
}
func (*Impl) Panicln(v ...any) {
	log.Panicln(v...)
}
func (*Impl) Prefix() string {
	return log.Prefix()
}
func (*Impl) Print(v ...any) {
	log.Print(v...)
}
func (*Impl) Printf(format string, v ...any) {
	log.Printf(format, v...)
}
func (*Impl) Println(v ...any) {
	log.Println(v...)
}
func (*Impl) SetFlags(flag int) {
	log.SetFlags(flag)
}
func (*Impl) SetOutput(w io.Writer) {
	log.SetOutput(w)
}
func (*Impl) SetPrefix(prefix string) {
	log.SetPrefix(prefix)
}
func (*Impl) Writer() io.Writer {
	return log.Writer()
}
