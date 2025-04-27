package opcodes

import (
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/utils"
)

type Load struct{}

var _ OpCode = (*Load)(nil)

func (op *Load) Apply(m machine.Machine) {
	reg := m.GetMemory()[m.GetIP()]
	m.IncrementIP()

	val := m.GetRegisters()[utils.ParseInt([]byte{reg})]
	m.Push(val)
}
