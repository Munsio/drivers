package lis3dh

// Constants/addresses used for I2C.

// The I2C addresses which this device listens to.
const (
	Address0 = 0x18 // SA0 is low
	Address1 = 0x19 // SA0 is high
)

// Registers. Names, addresses and comments copied from the datasheet.
const (
	WHO_AM_I = 0x0F
)

type DataRate uint8

// Data rate constants.
const (
	DataRate800Hz DataRate = iota // 800Hz,  1.25ms interval
	DataRate400Hz                 // 400Hz,  2.5ms  interval
	DataRate200Hz                 // 200Hz,  5ms    interval
	DataRate100Hz                 // 100Hz,  10ms   interval
	DataRate50Hz                  // 50Hz,   20ms   interval
	DataRate12Hz                  // 12.5Hz, 80ms   interval
	DataRate6Hz                   // 6.25Hz, 160ms  interval
	DataRate2Hz                   // 1.56Hz, 640ms  interval
)

type Sensitivity uint8

// Sensitivity constants.
const (
	Sensitivity2G Sensitivity = iota
	Sensitivity4G
	Sensitivity8G
)
