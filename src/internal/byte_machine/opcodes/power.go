package opcodes

import (
	"math"

	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
)

type Power struct{}

var _ OpCode = (*Power)(nil)

func (op *Power) Apply(m machine.Machine) {
	r := m.Pop()
	l := m.Pop()
	m.Push(int(math.Pow(float64(l), float64(r))))
}
