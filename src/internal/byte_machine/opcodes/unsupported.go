package opcodes

import (
	"fmt"

	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
)

type UnImplemented struct{}

var _ OpCode = (*Halt)(nil)

func (op *UnImplemented) Apply(m machine.Machine) {
	fmt.Printf("The opcode %v is not currently implemented. Please contact Kasey Abu-Rayyan for help in getting this added.\n", m.GetMemory()[m.GetIP()-1])
	m.Halt()
}
