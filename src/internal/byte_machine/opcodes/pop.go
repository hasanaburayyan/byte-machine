package opcodes

import (
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
)

type Pop struct{}

var _ OpCode = (*Pop)(nil)

func (op *Pop) Apply(m machine.Machine) {
	m.Pop()
}
