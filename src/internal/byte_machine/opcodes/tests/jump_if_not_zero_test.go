package opcodes_test

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestJumpIfNotZero_NonZeroValue_Jumps(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Memory: []byte{0x07}, // jump target 7
		Stack:  []int{42},    // top of stack is non-zero
		IP:     0,
	}

	jumpIfNotZero := &opcodes.JumpIfNotZero{}
	jumpIfNotZero.Apply(machine)

	require.Equal(t, 7, machine.GetIP())
}

func TestJumpIfNotZero_ZeroValue_NoJump(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Memory: []byte{0x07}, // jump target 7
		Stack:  []int{0},     // top of stack is zero
		IP:     0,
	}

	jumpIfNotZero := &opcodes.JumpIfNotZero{}
	jumpIfNotZero.Apply(machine)

	// Should not jump, IP just incremented by 1
	require.Equal(t, 1, machine.GetIP())
}
