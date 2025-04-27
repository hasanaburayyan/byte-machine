package opcodes

import (
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
)

type GreaterOrEq struct{}

var _ OpCode = (*GreaterOrEq)(nil)

func (op *GreaterOrEq) Apply(m machine.Machine) {
	r := m.Pop()
	l := m.Pop()

	if l >= r {
		m.Push(1)
		return
	}

	m.Push(0)
}
