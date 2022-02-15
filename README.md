# Facades

Provides mockable interfaces for (nearly) every package in the standard library.

[![Go Reference](https://pkg.go.dev/badge/jonwillia.ms/facades.svg)](https://pkg.go.dev/jonwillia.ms/facades)

## Notes

- Will only work with go1.18 on Linux/amd64 because the go.mod specifies 1.18 and there are likely undefined symbols when GOOS / GOARCH do not match.
I'll work build tag support in later and it should work across go version and targets.
- The tool for building this is here; it may be of general interest and can mock other packages: https://github.com/WIZARDISHUNGRY/facade
- I made this because once upon a time I wanted it -- I'm not sure I still do!