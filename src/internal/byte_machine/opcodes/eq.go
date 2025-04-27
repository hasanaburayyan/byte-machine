package opcodes

import (
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
)

type Eq struct{}

var _ OpCode = (*Eq)(nil)

func (op *Eq) Apply(m machine.Machine) {
	if m.Pop() == m.Pop() {
		m.Push(1)
		return
	}

	m.Push(0)
}
