package main

import (
	"fmt"
	"io"
	"os"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
)

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("No input detected. Exiting.")
		return
	}

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("failed to read input from stdin")
	}

	// Create a new ByteMachine
	machine := bytemachine.NewByteMachine(input)

	// Run the machine
	machine.Run()
}
