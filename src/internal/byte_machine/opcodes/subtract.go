package opcodes

import (
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
)

type Subtract struct{}

var _ OpCode = (*Subtract)(nil)

func (op *Subtract) Apply(m machine.Machine) {
	r := m.Pop()
	l := m.Pop()
	m.Push(l - r)
}
