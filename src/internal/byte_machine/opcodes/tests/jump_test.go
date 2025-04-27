package opcodes_test

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestJumpApply(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Memory: []byte{0x05}, // target address 5
		IP:     0,
	}

	jump := &opcodes.Jump{}
	jump.Apply(machine)

	require.Equal(t, 5, machine.GetIP())
}
