package opcodes

import (
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/utils"
)

type JumpIfZero struct{}

var _ OpCode = (*JumpIfZero)(nil)

func (op *JumpIfZero) Apply(m machine.Machine) {
	if m.Pop() == 0 {
		location := utils.ParseInt([]byte{m.GetMemory()[m.GetIP()]})
		m.SetIP(location)
		return
	}

	// Otherwise just increment over the jump location
	m.IncrementIP()
}
