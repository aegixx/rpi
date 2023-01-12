package monitor

import (
	"fmt"

	"github.com/aegixx/go-dht"
	"github.com/d2r2/go-logger"
	"github.com/stianeikeland/go-rpio/v4"
)

func Ready(buttonPin int, tempPin int) {
	// Open and initialize the library
	err := rpio.Open()
	if err != nil {
		panic(err)
	}
	defer rpio.Close()

	// Reset library log level to INFO
	logger.ChangePackageLogLevel("dht", logger.InfoLevel)

	button := rpio.Pin(buttonPin)
	button.Input()

	temp := rpio.Pin(tempPin)
	temp.Input()

	button.Detect(rpio.RiseEdge)

	// Keep checking for the button press
	for {
		if button.EdgeDetected() {

			// Read sensor
			tempC, humidity, _, err := dht.ReadDHTxxWithRetry(dht.DHT11, tempPin, false, 10)
			if err != nil {
				panic(err)
			}

			// Convert to Fahrenheit
			tempF := tempC*1.8 + 32

			// Show the user
			fmt.Printf("Temperature = %v°F, Humidity = %v%%\n", tempF, humidity)
		}
	}
}
