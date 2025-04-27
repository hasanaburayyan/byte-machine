package opcodes_test

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestLessApply(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{2, 5},
	}

	less := &opcodes.Less{}
	less.Apply(machine)

	require.Equal(t, []int{1}, machine.GetStack())
}

func TestLessApply_NotLess(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{6, 3},
	}

	less := &opcodes.Less{}
	less.Apply(machine)

	require.Equal(t, []int{0}, machine.GetStack())
}
