package gossd1305

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

func (v *SSD1305) InitDisplay(i2c *i2c.Options, width byte, height byte) {
	v = &SSD1305{}
	v.height = height
	v.width = width
	v.frame_buffer = make([]byte, (v.height/8)*v.width+1)
}

/*
func write_cmd(i2c *i2c.I2C, cmd byte) error {
	_, err := i2c.WriteBytes([]byte{0x80, cmd})
	if err != nil {
		fmt.Println(err)
		lg.Fatal(err)
		return err
	}
	return nil
}
*/
