package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aegixx/rpi/internal/app/button"
	"github.com/aegixx/rpi/internal/app/monitor"
)

func main() {

	args := os.Args[1:]

	buttonPin, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	ledPin, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}

	tempPin, err := strconv.Atoi(args[2])
	if err != nil {
		panic(err)
	}

	go button.Ready(buttonPin, ledPin)

	fmt.Printf("Press Button (pin %d) to measure temperature & humidity (pin %d)\n", buttonPin, tempPin)
	monitor.Ready(buttonPin, tempPin)
}
