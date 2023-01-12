package blinker

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func Blink(pin int, rate int) {
	// Open and initialize the library
	err := rpio.Open()
	if err != nil {
		panic(err)
	}
	defer rpio.Close()

	ledPin := rpio.Pin(pin)

	// Set the pin as an output
	ledPin.Output()

	// Blink X times per second
	period := 1000 / rate
	for {
		// Period is broken equally into HIGH and LOW
		fmt.Println("Blink")
		ledPin.High()
		time.Sleep(time.Duration(period/2) * time.Millisecond)
		ledPin.Low()
		time.Sleep(time.Duration(period/2) * time.Millisecond)
	}
}
