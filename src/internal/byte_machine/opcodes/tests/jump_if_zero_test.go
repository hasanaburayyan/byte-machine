package opcodes_test

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestJumpIfZero_ZeroValue_Jumps(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Memory: []byte{0x05}, // jump target 5
		Stack:  []int{0},     // top of stack is zero
		IP:     0,
	}

	jumpIfZero := &opcodes.JumpIfZero{}
	jumpIfZero.Apply(machine)

	require.Equal(t, 5, machine.GetIP())
}

func TestJumpIfZero_NonZeroValue_NoJump(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Memory: []byte{0x05}, // jump target 5
		Stack:  []int{1},     // top of stack is non-zero
		IP:     0,
	}

	jumpIfZero := &opcodes.JumpIfZero{}
	jumpIfZero.Apply(machine)

	// Should not jump, IP just incremented by 1
	require.Equal(t, 1, machine.GetIP())
}
