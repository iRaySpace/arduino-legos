package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/servo"
)

func main() {
	s, err := servo.New(machine.Timer1, machine.D9)
	if err != nil {
		for {
			println(err)
			time.Sleep(time.Second)
		}
		return
	}

	for {
		s.SetAngle(0)
		time.Sleep(1 * time.Second)

		s.SetAngle(180)
		time.Sleep(1 * time.Second)
	}
}
