package opcodes_test

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestPushApply(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Memory: []byte{0x05},
		IP:     0,
	}

	push := &opcodes.Push{}
	push.Apply(machine)

	require.Equal(t, []int{5}, machine.GetStack())
}
