package debugger

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
)

// Debugger provides debugging capabilities for a ByteMachine.
type Debugger struct {
	machine     *bytemachine.ByteMachine
	breakpoints map[int]bool // Maps IP addresses to breakpoint status
	running     bool         // Indicates if the debugger is running
}

// NewDebugger creates a new Debugger for the given ByteMachine.
func NewDebugger(m *bytemachine.ByteMachine) *Debugger {
	return &Debugger{
		machine:     m,
		breakpoints: make(map[int]bool),
		running:     true,
	}
}

// Run starts the debugger's interactive loop.
func (d *Debugger) Run() {
	fmt.Println("ByteMachine Debugger. Commands: step, continue, break <ip>, state, quit")
	scanner := bufio.NewScanner(os.Stdin)

	for d.running && !d.machine.Halted {
		d.printNextInstruction()
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		d.handleCommand(input)
	}

	if d.machine.Halted {
		fmt.Println("Program halted.")
		d.printState()
	}
}

// handleCommand processes user input commands.
func (d *Debugger) handleCommand(input string) {
	args := strings.Fields(input)
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "step", "s":
		d.step()
	case "continue", "c":
		d.continueExecution()
	case "break", "b":
		if len(args) != 2 {
			fmt.Println("Usage: break <ip>")
			return
		}
		ip, err := strconv.Atoi(args[1])
		if err != nil || ip < 0 || ip >= len(d.machine.Memory) {
			fmt.Println("Invalid IP address")
			return
		}
		d.breakpoints[ip] = true
		fmt.Printf("Breakpoint set at IP %d\n", ip)
	case "state", "st":
		d.printState()
	case "quit", "q":
		d.running = false
		fmt.Println("Exiting debugger")
	default:
		fmt.Println("Unknown command. Available: step, continue, break <ip>, state, quit")
	}
}

// step executes a single instruction.
func (d *Debugger) step() {
	if d.machine.Halted {
		fmt.Println("Program has halted")
		return
	}
	if d.machine.IP >= len(d.machine.Memory) {
		fmt.Println("Reached end of program")
		d.machine.Halted = true
		return
	}

	opcode := d.machine.Memory[d.machine.IP]
	d.machine.IP++

	if op, ok := opcodes.OpCodeDefs[opcode]; ok {
		op.Op.Apply(d.machine)
	} else {
		fmt.Printf("Unknown opcode: %v\n", opcode)
		d.machine.Halted = true
	}
}

// continueExecution runs the program until a breakpoint or halt.
func (d *Debugger) continueExecution() {
	for !d.machine.Halted && d.machine.IP < len(d.machine.Memory) {
		if d.breakpoints[d.machine.IP] {
			fmt.Printf("Hit breakpoint at IP %d\n", d.machine.IP)
			return
		}
		d.step()
	}
}

// printState displays the current state of the ByteMachine.
func (d *Debugger) printState() {
	fmt.Printf("IP: %d\n", d.machine.IP)
	fmt.Printf("Stack: %v\n", d.machine.Stack)
	fmt.Printf("Registers: %v\n", d.machine.Registers)
	fmt.Printf("Halted: %v\n", d.machine.Halted)
	fmt.Printf("Memory: %v\n", d.machine.Memory)
	fmt.Println("Breakpoints:", d.breakpoints)
}

// printNextInstruction displays the instruction at the current IP.
func (d *Debugger) printNextInstruction() {
	if d.machine.IP >= len(d.machine.Memory) {
		fmt.Println("No more instructions")
		return
	}
	opcode := d.machine.Memory[d.machine.IP]
	if op, ok := opcodes.OpCodeDefs[opcode]; ok {
		fmt.Printf("Next instruction (IP %d): %s\n", d.machine.IP, op.Name)
	} else {
		fmt.Printf("Next instruction (IP %d): Unknown opcode %v\n", d.machine.IP, opcode)
	}
}
