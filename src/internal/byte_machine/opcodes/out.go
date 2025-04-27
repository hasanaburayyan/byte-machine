package opcodes

import (
	"fmt"

	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"
)

type Out struct{}

var _ OpCode = (*Push)(nil)

func (op *Out) Apply(m machine.Machine) {
	fmt.Printf("%v\n", m.Peek())
}
