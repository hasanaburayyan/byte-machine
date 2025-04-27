package opcodes

import "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"

type NOOP struct{}

var _ OpCode = (*NOOP)(nil)

func (op *NOOP) Apply(bm machine.Machine) { /** No op! */ }
