package opcodes_test

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestPopApply(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{5},
	}

	pop := &opcodes.Pop{}
	pop.Apply(machine)

	require.Empty(t, machine.GetStack())
}
