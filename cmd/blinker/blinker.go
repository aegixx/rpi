package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aegixx/rpi/internal/app/blinker"
)

func main() {

	args := os.Args[1:]

	pin, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	rate, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}

	fmt.Println("Blinking LED on pin", pin, "at a rate of", rate, "per second")
	blinker.Blink(pin, rate)
}
