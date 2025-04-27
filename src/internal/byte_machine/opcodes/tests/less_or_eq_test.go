package opcodes_test

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestLessOrEqApply_Less(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{3, 5},
	}

	lessOrEq := &opcodes.LessOrEq{}
	lessOrEq.Apply(machine)

	require.Equal(t, []int{1}, machine.GetStack())
}

func TestLessOrEqApply_Equal(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{4, 4},
	}

	lessOrEq := &opcodes.LessOrEq{}
	lessOrEq.Apply(machine)

	require.Equal(t, []int{1}, machine.GetStack())
}

func TestLessOrEqApply_NotLess(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{6, 5},
	}

	lessOrEq := &opcodes.LessOrEq{}
	lessOrEq.Apply(machine)

	require.Equal(t, []int{0}, machine.GetStack())
}
