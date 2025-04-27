package opcodes

import (
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/utils"
)

type Jump struct{}

var _ OpCode = (*Jump)(nil)

func (op *Jump) Apply(m machine.Machine) {
	location := utils.ParseInt([]byte{m.GetMemory()[m.GetIP()]})
	m.SetIP(location)
}
