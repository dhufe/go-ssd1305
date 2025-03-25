package main

import (
	i2c "github.com/googolgl/go-i2c"
)

func main() {
	i2cDevice, err := i2c.New(0x3c, "/dev/i2c-1")
	if err != nil {
		i2cDevice.Log.Fatal(err)
	}
	// Free I2C connection on exit
	defer i2cDevice.Close()

	// Set log level: 0 - Panic, 1 - Fatal, 2 - Error, 3 - Warning, 4 - Info, 5 - Debug
	i2cDevice.Log.SetLevel(5)

	_, err = InitDisplay(i2cDevice, 128, 32)
}
