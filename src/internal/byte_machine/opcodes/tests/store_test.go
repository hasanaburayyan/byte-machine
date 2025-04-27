package opcodes_test

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestStoreApply(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Stack:  []int{123},
		Memory: []byte{0x02}, // Store into Register[2]
		IP:     0,
	}

	store := &opcodes.Store{}
	store.Apply(machine)

	require.Equal(t, 123, machine.GetRegisters()[2])
}
