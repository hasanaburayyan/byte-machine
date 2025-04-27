package opcodes

import (
	"testing"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	machine := &bytemachine.ByteMachine{
		Memory: []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07},
		Registers: [8]int{
			100, 200, 300, 400, 500, 600, 700, 800,
		},
		IP:    0,
		Stack: []int{},
	}

	load := &opcodes.Load{}

	for _, regVale := range machine.Registers {
		load.Apply(machine)
		require.Equal(t, regVale, machine.Pop())
	}
}
