package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aegixx/rpi/internal/app/button"
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

	fmt.Printf("Press Button (pin %d) to light LED (pin %d)\n", buttonPin, ledPin)
	button.Ready(buttonPin, ledPin)
}
