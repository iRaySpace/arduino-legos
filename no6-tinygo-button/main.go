package main

import (
	"machine"
	"time"
)

func momentaryCircuit(btn *machine.Pin, led *machine.Pin) {
	for {
		btnState := btn.Get() // Pullup is on a "High" state
		if !btnState {
			led.High()
		} else {
			led.Low()
		}
	}
}

func toggleCircuit(btn *machine.Pin, led *machine.Pin) {
	lastTimeChanged := time.Now().UnixMilli()
	lastBtnState := btn.Get() // Pullup default
	for {
		btnState := btn.Get()
		currentTime := time.Now().UnixMilli()
		if currentTime-lastTimeChanged > 20 { // debounce
			if !btnState && lastBtnState {
				lastTimeChanged = time.Now().UnixMilli()
				led.Set(!led.Get())
			}
		}
		lastBtnState = btnState
	}
}

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led.Low()

	btn := machine.D7
	btn.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	// momentaryCircuit(&btn, &led)
	toggleCircuit(&btn, &led)
}
