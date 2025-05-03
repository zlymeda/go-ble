package dev

import (
	"github.com/zlymeda/go-ble"
	"github.com/zlymeda/go-ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
