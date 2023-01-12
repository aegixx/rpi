package button

import (
	"fmt"

	"github.com/stianeikeland/go-rpio/v4"
)

func Ready(buttonPin int, ledPin int) {
	// Open and initialize the library
	err := rpio.Open()
	if err != nil {
		panic(err)
	}
	defer rpio.Close()

	led := rpio.Pin(ledPin)
	led.Output()

	button := rpio.Pin(buttonPin)
	button.Input()
	button.Detect(rpio.RiseEdge)

	// Keep checking for the button press
	for {
		if button.EdgeDetected() {
			fmt.Println("Button pressed")
		}

		if button.Read() == rpio.Low {
			led.High()
		} else {
			led.Low()
		}
	}
}
