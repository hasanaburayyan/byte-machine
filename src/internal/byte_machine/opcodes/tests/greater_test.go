package opcodes_test

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestGreaterApply(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{7, 3},
	}

	greater := &opcodes.Greater{}
	greater.Apply(machine)

	require.Equal(t, []int{1}, machine.GetStack())
}

func TestGreaterApply_NotGreater(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{2, 5},
	}

	greater := &opcodes.Greater{}
	greater.Apply(machine)

	require.Equal(t, []int{0}, machine.GetStack())
}
