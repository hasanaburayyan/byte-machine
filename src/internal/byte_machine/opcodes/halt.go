package opcodes

import (
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
)

type Halt struct{}

var _ OpCode = (*Halt)(nil)

func (op *Halt) Apply(m machine.Machine) {
	m.Halt()
}
