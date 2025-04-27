package opcodes

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestHalt(t *testing.T) {
	machine := bytemachine.NewByteMachine([]byte{})
	halt := &opcodes.Halt{}

	halt.Apply(machine)

	require.Equal(t, true, machine.Halted)
}
