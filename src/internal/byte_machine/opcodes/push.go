package opcodes

import (
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/utils"
)

type Push struct{}

var _ OpCode = (*Push)(nil)

func (op *Push) Apply(m machine.Machine) {
	val := m.GetMemory()[m.GetIP()]
	m.Push(utils.ParseInt([]byte{val}))
	m.IncrementIP()
}
