package opcodes_test

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestNeApply(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{5, 3},
	}

	ne := &opcodes.Ne{}
	ne.Apply(machine)

	require.Equal(t, []int{1}, machine.GetStack())
}

func TestNeApply_Equal(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{4, 4},
	}

	ne := &opcodes.Ne{}
	ne.Apply(machine)

	require.Equal(t, []int{0}, machine.GetStack())
}
