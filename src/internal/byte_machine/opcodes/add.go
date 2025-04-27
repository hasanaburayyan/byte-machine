package opcodes

import (
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
)

type Add struct{}

var _ OpCode = (*Add)(nil)

func (op *Add) Apply(m machine.Machine) {
	m.Push(m.Pop() + m.Pop())
}
