package bytemachine

import (
	"fmt"

	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
)

type ByteMachine struct {
	Memory    []byte
	IP        int
	Stack     []int
	Registers [8]int
	Halted    bool
}

var _ machine.Machine = (*ByteMachine)(nil)

func NewByteMachine(program []byte) *ByteMachine {
	return &ByteMachine{
		Memory: program,
		IP:     0,
	}
}

func (bm *ByteMachine) Run() int {
	for bm.IP < len(bm.Memory) {
		opcode := bm.Memory[bm.IP]
		bm.IP++

		if op, ok := opcodes.OpCodeDefs[opcode]; ok {
			op.Op.Apply(bm)
		} else {
			fmt.Printf("unknown opcode: %v\n", opcode)
			bm.Halted = true
		}

		if bm.Halted {
			break
		}
	}

	return 0
}

func (bm *ByteMachine) GetMemory() []byte {
	return bm.Memory
}

func (bm *ByteMachine) GetIP() int {
	return bm.IP
}

func (bm *ByteMachine) IncrementIP() {
	bm.IP++
}

func (bm *ByteMachine) DecrementIP() {
	bm.IP--
}

func (bm *ByteMachine) SetIP(ip int) {
	bm.IP = ip
}

func (bm *ByteMachine) GetStack() []int {
	return bm.Stack
}

func (bm *ByteMachine) Push(val int) {
	bm.Stack = append(bm.Stack, val)
}

func (bm *ByteMachine) Pop() int {
	if len(bm.Stack) == 0 {
		panic("Pop from empty stack")
	}

	// Get top value
	val := bm.Stack[len(bm.Stack)-1]

	// Remove it
	bm.Stack = bm.Stack[:len(bm.Stack)-1]

	return val
}

func (bm *ByteMachine) Peek() int {
	// Get top value
	val := bm.Stack[len(bm.Stack)-1]
	return val
}

func (bm *ByteMachine) GetRegisters() [8]int {
	return bm.Registers
}

func (bm *ByteMachine) SetRegister(reg int, val int) {
	bm.Registers[reg] = val
}

func (bm *ByteMachine) GetHalted() bool {
	return bm.Halted
}

func (bm *ByteMachine) Halt() {
	bm.Halted = true
}
