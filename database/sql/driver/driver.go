// Code generated by a tool. DO NOT EDIT.

// Package driver provides a mockable wrapper for database/sql/driver.
package driver

import (
	driver "database/sql/driver"
)

var _ Interface = &Impl{}
var _ = driver.IsScanValue

type Interface interface {
	IsScanValue(v any) bool
	IsValue(v any) bool
}

type Impl struct{}

func (*Impl) IsScanValue(v any) bool {
	return driver.IsScanValue(v)
}
func (*Impl) IsValue(v any) bool {
	return driver.IsValue(v)
}
