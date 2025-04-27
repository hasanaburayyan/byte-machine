package opcodes_test

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestSubtractApply(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{7, 2},
	}

	subtract := &opcodes.Subtract{}
	subtract.Apply(machine)

	require.Equal(t, []int{5}, machine.GetStack())
}
