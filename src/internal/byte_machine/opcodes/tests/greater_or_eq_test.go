package opcodes_test

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestGreaterOrEqApply_Greater(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{7, 3},
	}

	greaterOrEq := &opcodes.GreaterOrEq{}
	greaterOrEq.Apply(machine)

	require.Equal(t, []int{1}, machine.GetStack())
}

func TestGreaterOrEqApply_Equal(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{5, 5},
	}

	greaterOrEq := &opcodes.GreaterOrEq{}
	greaterOrEq.Apply(machine)

	require.Equal(t, []int{1}, machine.GetStack())
}

func TestGreaterOrEqApply_NotGreater(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{2, 5},
	}

	greaterOrEq := &opcodes.GreaterOrEq{}
	greaterOrEq.Apply(machine)

	require.Equal(t, []int{0}, machine.GetStack())
}
