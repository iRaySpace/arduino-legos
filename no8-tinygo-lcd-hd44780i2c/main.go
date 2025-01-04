package main

import (
	"machine"
	"strconv"
	"time"

	"tinygo.org/x/drivers/hd44780i2c"
)

func main() {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: 400e3, // 31-400kHz/400e3
	})

	lcd := hd44780i2c.New(machine.I2C0, 0) // by default uses 0x27
	lcd.Configure(hd44780i2c.Config{
		Width:       16,
		Height:      2,
		CursorOn:    true,
		CursorBlink: true,
	})

	lcd.Print([]byte("Henlo!"))

	lcd.CreateCharacter(0x0, []byte{0x00, 0x11, 0x0E, 0x1F, 0x15, 0x1F, 0x1F, 0x1F})
	lcd.Print([]byte{0x0})

	time.Sleep(time.Millisecond * 7000)

	lcd.CursorOn(false)
	lcd.CursorBlink(false)

	i := 0
	for {
		lcd.ClearDisplay()
		lcd.SetCursor(0, 0)
		lcd.Print([]byte(strconv.FormatInt(int64(i), 10)))
		i++
		time.Sleep(time.Millisecond * 100)
	}
}
