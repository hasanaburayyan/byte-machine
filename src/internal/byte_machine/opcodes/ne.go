package opcodes

import (
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
)

type Ne struct{}

var _ OpCode = (*Ne)(nil)

func (op *Ne) Apply(m machine.Machine) {
	if m.Pop() != m.Pop() {
		m.Push(1)
		return
	}

	m.Push(0)
}
