package main

import (
	"fmt"
	"os"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
)

func main() {
	// Read program.bin
	program, err := os.ReadFile("program.bin")
	if err != nil {
		fmt.Printf("Failed to read program.bin: %v\n", err)
		os.Exit(1)
	}

	// Create a new ByteMachine
	machine := bytemachine.NewByteMachine(program)

	// Run the machine
	machine.Run()
}
