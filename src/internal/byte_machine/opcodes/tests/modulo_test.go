package opcodes_test

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestModuloApply(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{9, 4},
	}

	modulo := &opcodes.Modulo{}
	modulo.Apply(machine)

	require.Equal(t, []int{1}, machine.GetStack())
}
