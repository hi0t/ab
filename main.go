package main

import (
	"fmt"

	"github.com/hi0t/ab/pi"
)

func main() {
	pi.SetupGpio()

	pi.OpenPin(4)

	fmt.Println("Hello world!")
}
