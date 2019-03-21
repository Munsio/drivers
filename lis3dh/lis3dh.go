// Package lis3dh provides a driver for the LIS3DH 3-axis accelerometer by
// STMicroelectronics.
//
// Datasheet:
// https://www.st.com/resource/en/datasheet/cd00274221.pdf
package lis3dh

import (
	"machine"
)

// Device wraps an I2C connection to a LIS3DH device.
type Device struct {
	bus         machine.I2C
	address     uint8
	sensitivity Sensitivity
}

// New creates a new LIS3DH connection. The I2C bus must already be configured.
//
// This function only creates the Device object, it does not touch the device.
func New(bus machine.I2C, address uint8) Device {
	return Device{bus, address, Sensitivity2G}
}

// Connected returns whether a LIS3DH has been found.
// It does a "who am I" request and checks the response.
func (d Device) Connected() bool {
	data := []byte{0}
	err := d.bus.ReadRegister(d.address, WHO_AM_I, data)
	if err != nil {
		println(err.Error())
	}
	println("whoami:", data[0])
	return data[0] == 0x33
}
