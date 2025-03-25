package main

import (
	i2c "github.com/googolgl/go-i2c"
)

const (
	SET_CONTRAST        = 0x81
	SET_ENTIRE_ON       = 0xA4
	SET_NORM_INV        = 0xA6
	SET_DISP            = 0xAE
	SET_MEM_ADDR        = 0x20
	SET_COL_ADDR        = 0x21
	SET_PAGE_ADDR       = 0x22
	SET_DISP_START_LINE = 0x40
	SET_LUT             = 0x91
	SET_SEG_REMAP       = 0xA0
	SET_MUX_RATIO       = 0xA8
	SET_MASTER_CONFIG   = 0xAD
	SET_COM_OUT_DIR     = 0xC0
	SET_COMSCAN_DEC     = 0xC8
	SET_DISP_OFFSET     = 0xD3
	SET_COM_PIN_CFG     = 0xDA
	SET_DISP_CLK_DIV    = 0xD5
	SET_AREA_COLOR      = 0xD8
	SET_PRECHARGE       = 0xD9
	SET_VCOM_DESEL      = 0xDB
	SET_CHARGE_PUMP     = 0x8D
)

type SSD1305 struct {
	width        byte
	height       byte
	frame_buffer []byte
}

// InitDisplay prepared the internal struct and the framebuffer of the display
func (v *SSD1305) InitDisplay(i2c *i2c.Options, width byte, height byte) {
	v = &SSD1305{}
	v.height = height
	v.width = width
	v.frame_buffer = make([]byte, (v.height/8)*v.width+1)

	_init(i2c, v)
}

// private function for sending commands to the display
func write_cmd(i2c *i2c.Options, cmd byte) error {
	_, err := i2c.WriteBytes([]byte{0x80, cmd})
	if err != nil {
		return err
	}
	return nil
}

func Write_scr(i2c *i2c.Options, v *SSD1305) error {
	err := write_cmd(i2c, SET_COL_ADDR)
	err = write_cmd(i2c, 4)
	err = write_cmd(i2c, 131)
	err = write_cmd(i2c, SET_PAGE_ADDR)
	err = write_cmd(i2c, 0)
	err = write_cmd(i2c, 3)
	_, err = i2c.WriteBytes(v.frame_buffer)
	return err
}

func _init(i2c *i2c.Options, v *SSD1305) error {
	write_cmd(i2c, SET_DISP|0x00)
	write_cmd(i2c, SET_DISP_CLK_DIV)
	write_cmd(i2c, 0x80)
	write_cmd(i2c, SET_SEG_REMAP|0x01)
	write_cmd(i2c, SET_MUX_RATIO)
	write_cmd(i2c, 0x1F)
	write_cmd(i2c, SET_DISP_OFFSET)
	write_cmd(i2c, 0x00)
	write_cmd(i2c, SET_MASTER_CONFIG)
	write_cmd(i2c, 0x8E)
	write_cmd(i2c, SET_AREA_COLOR)
	write_cmd(i2c, 0x05)
	write_cmd(i2c, SET_MEM_ADDR)
	write_cmd(i2c, 0x00)
	write_cmd(i2c, SET_DISP_START_LINE)
	write_cmd(i2c, 0x2E)
	write_cmd(i2c, SET_COMSCAN_DEC)
	write_cmd(i2c, SET_COM_PIN_CFG)
	write_cmd(i2c, 0x12)
	write_cmd(i2c, SET_LUT)
	write_cmd(i2c, 0x3F)
	write_cmd(i2c, 0x3F)
	write_cmd(i2c, 0x3F)
	write_cmd(i2c, 0x3F)
	write_cmd(i2c, SET_CONTRAST)
	write_cmd(i2c, 0xFF)
	write_cmd(i2c, SET_PRECHARGE)
	write_cmd(i2c, 0xD2)
	write_cmd(i2c, SET_VCOM_DESEL)
	write_cmd(i2c, 0x34)
	write_cmd(i2c, SET_NORM_INV)
	write_cmd(i2c, SET_ENTIRE_ON)
	write_cmd(i2c, SET_CHARGE_PUMP)
	write_cmd(i2c, 0x14)
	write_cmd(i2c, SET_DISP|0x01)

	return nil
}
