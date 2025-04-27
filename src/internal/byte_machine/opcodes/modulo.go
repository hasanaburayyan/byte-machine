package opcodes

import (
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
)

type Modulo struct{}

var _ OpCode = (*Modulo)(nil)

func (op *Modulo) Apply(m machine.Machine) {
	r := m.Pop()
	l := m.Pop()
	m.Push(l % r)
}
