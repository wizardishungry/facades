// Code generated by a tool. DO NOT EDIT.

// Package time provides a mockable wrapper for time.
package time

import (
	time "time"
)

var _ Interface = &Impl{}
var _ = time.After

type Interface interface {
	After(d time.Duration) <-chan time.Time
	AfterFunc(d time.Duration, f func()) *time.Timer
	Date(year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location) time.Time
	FixedZone(name string, offset int) *time.Location
	LoadLocation(name string) (*time.Location, error)
	LoadLocationFromTZData(name string, data []byte) (*time.Location, error)
	NewTicker(d time.Duration) *time.Ticker
	NewTimer(d time.Duration) *time.Timer
	Now() time.Time
	Parse(layout string, value string) (time.Time, error)
	ParseDuration(s string) (time.Duration, error)
	ParseInLocation(layout string, value string, loc *time.Location) (time.Time, error)
	Since(t time.Time) time.Duration
	Sleep(d time.Duration)
	Tick(d time.Duration) <-chan time.Time
	Unix(sec int64, nsec int64) time.Time
	UnixMicro(usec int64) time.Time
	UnixMilli(msec int64) time.Time
	Until(t time.Time) time.Duration
}

type Impl struct{}

func (*Impl) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}
func (*Impl) AfterFunc(d time.Duration, f func()) *time.Timer {
	return time.AfterFunc(d, f)
}
func (*Impl) Date(year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location) time.Time {
	return time.Date(year, month, day, hour, min, sec, nsec, loc)
}
func (*Impl) FixedZone(name string, offset int) *time.Location {
	return time.FixedZone(name, offset)
}
func (*Impl) LoadLocation(name string) (*time.Location, error) {
	return time.LoadLocation(name)
}
func (*Impl) LoadLocationFromTZData(name string, data []byte) (*time.Location, error) {
	return time.LoadLocationFromTZData(name, data)
}
func (*Impl) NewTicker(d time.Duration) *time.Ticker {
	return time.NewTicker(d)
}
func (*Impl) NewTimer(d time.Duration) *time.Timer {
	return time.NewTimer(d)
}
func (*Impl) Now() time.Time {
	return time.Now()
}
func (*Impl) Parse(layout string, value string) (time.Time, error) {
	return time.Parse(layout, value)
}
func (*Impl) ParseDuration(s string) (time.Duration, error) {
	return time.ParseDuration(s)
}
func (*Impl) ParseInLocation(layout string, value string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(layout, value, loc)
}
func (*Impl) Since(t time.Time) time.Duration {
	return time.Since(t)
}
func (*Impl) Sleep(d time.Duration) {
	time.Sleep(d)
}
func (*Impl) Tick(d time.Duration) <-chan time.Time {
	return time.Tick(d)
}
func (*Impl) Unix(sec int64, nsec int64) time.Time {
	return time.Unix(sec, nsec)
}
func (*Impl) UnixMicro(usec int64) time.Time {
	return time.UnixMicro(usec)
}
func (*Impl) UnixMilli(msec int64) time.Time {
	return time.UnixMilli(msec)
}
func (*Impl) Until(t time.Time) time.Duration {
	return time.Until(t)
}