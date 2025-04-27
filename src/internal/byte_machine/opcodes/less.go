package opcodes

import (
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
)

type Less struct{}

var _ OpCode = (*Less)(nil)

func (op *Less) Apply(m machine.Machine) {
	r := m.Pop()
	l := m.Pop()

	if l < r {
		m.Push(1)
		return
	}

	m.Push(0)
}
