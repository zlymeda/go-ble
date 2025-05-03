package dev

import (
	"github.com/zlymeda/go-ble"
	"github.com/zlymeda/go-ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
