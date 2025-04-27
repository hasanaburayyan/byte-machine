package opcodes

import (
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
)

type Multiply struct{}

var _ OpCode = (*Multiply)(nil)

func (op *Multiply) Apply(m machine.Machine) {
	m.Push(m.Pop() * m.Pop())
}
