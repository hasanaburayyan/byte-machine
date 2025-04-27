package opcodes

import (
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
)

type Divide struct{}

var _ OpCode = (*Divide)(nil)

func (op *Divide) Apply(m machine.Machine) {
	r := m.Pop()
	l := m.Pop()
	m.Push(l / r)
}
