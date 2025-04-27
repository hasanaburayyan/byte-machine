package opcodes_test

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestPowerApply(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack: []int{2, 3},
	}

	power := &opcodes.Power{}
	power.Apply(machine)

	require.Equal(t, []int{8}, machine.GetStack())
}
