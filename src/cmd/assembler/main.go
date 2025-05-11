package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hasanaburayyan/byte-machine/src/internal/assembler"
	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: byteasm <input.bm> [--run]")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	runMode := len(os.Args) >= 3 && os.Args[2] == "--run"

	// Open source file
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open input file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	// Assemble
	output, err := assembler.Assemble(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to assemble: %v\n", err)
		os.Exit(1)
	}

	if runMode {
		// Run it immediately
		machine := bytemachine.NewByteMachine(output)
		machine.Run()
	} else {
		// Write to .bin file
		outputFile := changeExtension(inputFile, ".bin")
		err = os.WriteFile(outputFile, output, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write output file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Assembled successfully to %s\n", outputFile)
	}
}

func changeExtension(path, newExt string) string {
	return filepath.Join(
		filepath.Dir(path),
		filepath.Base(path[:len(path)-len(filepath.Ext(path))])+newExt,
	)
}
