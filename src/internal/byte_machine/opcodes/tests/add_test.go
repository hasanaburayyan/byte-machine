package opcodes_test

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestAddApply(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{2, 3},
	}

	add := &opcodes.Add{}
	add.Apply(machine)

	require.Equal(t, []int{5}, machine.GetStack())
}
