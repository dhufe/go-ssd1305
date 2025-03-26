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

	fbuf, err := InitDisplay(i2cDevice, 128, 32)

	for p, i := range fbuf.frame_buffer {
		if p == 0 {
			fbuf.frame_buffer[i] = 0x40
		} else {
			fbuf.frame_buffer[i] = 0xff
		}
	}

	for x := 60; x < 80; x++ {
		for y := 5; y < 25; y++ {
			SetPixel(x, y, false, fbuf)
		}
	}

	err = Write_scr(i2cDevice, fbuf)
}
