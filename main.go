package main

import (
	"fmt"
	"os"
)

func main() {
	// x_direction := 1
	// y_direction := 0

	var handle *os.File

	if len(os.Args) >= 2 {
		handle, _ = os.Open(os.Args[1])
	} else {
		handle = os.Stdin
	}
	defer handle.Close()

	grid := Tokenize(handle)

	fmt.Println(grid)
}
