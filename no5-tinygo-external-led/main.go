package main

import (
	"machine"
	"time"
)

func main() {
	d8 := machine.D8
	d8.Configure(machine.PinConfig{Mode: machine.PinOutput})

	d2 := machine.D2
	d2.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		d8.High()
		d2.Low()
		time.Sleep(time.Millisecond * 1000)

		d8.Low()
		d2.High()
		time.Sleep(time.Millisecond * 1000)
	}
}
