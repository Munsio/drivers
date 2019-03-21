// Connects to a LIS3DH I2C accelerometer.
package main

import (
	"machine"
	"time"

	"github.com/tinygo-org/drivers/lis3dh"
)

var i2c = machine.I2C1

func main() {
	i2c.Configure(machine.I2CConfig{
		SDA: machine.SDA1_PIN,
		SCL: machine.SCL1_PIN,
	})

	accel := lis3dh.New(i2c, lis3dh.Address1) // address on the Circuit Playground Express
	accel.Configure(lis3dh.DataRate200Hz, mma8653.Sensitivity2G)

	for {
		x, y, z, _ := accel.ReadAcceleration()
		println(x, y, z)
		time.Sleep(time.Millisecond * 100)
	}
}
