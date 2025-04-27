package opcodes_test

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestEqApply(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{5, 5},
	}

	eq := &opcodes.Eq{}
	eq.Apply(machine)

	require.Equal(t, []int{1}, machine.GetStack())
}

func TestEqApply_NotEqual(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{5, 3},
	}

	eq := &opcodes.Eq{}
	eq.Apply(machine)

	require.Equal(t, []int{0}, machine.GetStack())
}
