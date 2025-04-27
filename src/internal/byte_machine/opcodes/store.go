package opcodes

import (
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/utils"
)

type Store struct{}

var _ OpCode = (*Store)(nil)

func (op *Store) Apply(m machine.Machine) {
	reg := m.GetMemory()[m.GetIP()]
	m.IncrementIP()

	val := m.Pop()
	m.SetRegister(utils.ParseInt([]byte{reg}), val)
}
