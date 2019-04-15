// Package lis3dh provides a driver for the LIS3DH digital accelerometer.
//
// Datasheet: https://www.st.com/resource/en/datasheet/lis3dh.pdf
package lis3dh

import (
	"machine"
	"time"
)

// Device wraps an I2C connection to a LIS3DH device.
type Device struct {
	bus     machine.I2C
	Address uint16
}

// New creates a new LIS3DH connection. The I2C bus must already be configured.
//
// This function only creates the Device object, it does not touch the device.
func New(bus machine.I2C) Device {
	return Device{bus: bus, Address: Address0}
}

// Configure sets up the device for communication
func (d *Device) Configure() {
	d.bus.WriteRegister(uint8(d.Address), REG_CTRL5, []byte{0x80})
	time.Sleep(5 * time.Millisecond) // reset takes 5ms

	// enable all axes, normal mode
	d.bus.WriteRegister(uint8(d.Address), REG_CTRL1, []byte{0x07})

	// 400Hz rate
	d.SetDataRate(DATARATE_400_HZ)

	// High res & BDU enabled
	d.bus.WriteRegister(uint8(d.Address), REG_CTRL4, []byte{0x88})

	// enable ADC
	d.bus.WriteRegister(uint8(d.Address), REG_TEMPCFG, []byte{0x80})
}

// Connected returns whether a LIS3DH has been found.
// It does a "who am I" request and checks the response.
func (d *Device) Connected() bool {
	data := []byte{0}
	err := d.bus.ReadRegister(uint8(d.Address), WHO_AM_I, data)
	if err != nil {
		return false
	}
	return data[0] == 0x33
}

// SetDataRate sets the speed of data collected by the LIS3DH.
func (d *Device) SetDataRate(rate DataRate) {
	ctl1 := []byte{0}
	err := d.bus.ReadRegister(uint8(d.Address), REG_CTRL1, ctl1)
	if err != nil {
		println(err.Error())
	}
	// mask off bits
	ctl1[0] &^= 0xf0
	ctl1[0] |= (byte(rate) << 4)
	d.bus.WriteRegister(uint8(d.Address), REG_CTRL1, ctl1)
}

// SetRange sets the G range for LIS3DH.
func (d *Device) SetRange(r Range) {
	ctl := []byte{0}
	err := d.bus.ReadRegister(uint8(d.Address), REG_CTRL4, ctl)
	if err != nil {
		println(err.Error())
	}
	// mask off bits
	ctl[0] &^= 0x30
	ctl[0] |= (byte(r) << 4)
	d.bus.WriteRegister(uint8(d.Address), REG_CTRL4, ctl)
}

// GetRange returns the current G range for LIS3DH.
func (d *Device) GetRange() (r Range) {
	ctl := []byte{0}
	err := d.bus.ReadRegister(uint8(d.Address), REG_CTRL4, ctl)
	if err != nil {
		println(err.Error())
	}
	// mask off bits
	r = Range(ctl[0] >> 4)
	r &= 0x03

	return r
}

// ReadAcceleration returns the adjusted x, y and z axis in milli-Gs.
func (d *Device) ReadAcceleration() (x int16, y int16, z int16) {
	x, y, z = d.ReadRawAcceleration()
	r := d.GetRange()
	divider := float32(1)
	switch r {
	case RANGE_16_G:
		divider = 1365
	case RANGE_8_G:
		divider = 4096
	case RANGE_4_G:
		divider = 8190
	case RANGE_2_G:
		divider = 16380
	}

	return int16(float32(x) / divider * 1000), int16(float32(y) / divider * 1000), int16(float32(z) / divider * 1000)
}

// ReadRawAcceleration returns the raw x, y and z axis from the LIS3DH
func (d *Device) ReadRawAcceleration() (x int16, y int16, z int16) {
	d.bus.WriteRegister(uint8(d.Address), REG_OUT_X_L|0x80, nil)

	data := []byte{0, 0, 0, 0, 0, 0}
	d.bus.Tx(d.Address, nil, data)

	x = int16((uint16(data[0]) << 8) | uint16(data[1]))
	y = int16((uint16(data[2]) << 8) | uint16(data[3]))
	z = int16((uint16(data[4]) << 8) | uint16(data[5]))

	return
}
