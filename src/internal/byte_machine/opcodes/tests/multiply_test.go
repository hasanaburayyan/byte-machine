package opcodes_test

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestMultiplyApply(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{3, 4},
	}

	multiply := &opcodes.Multiply{}
	multiply.Apply(machine)

	require.Equal(t, []int{12}, machine.GetStack())
}
