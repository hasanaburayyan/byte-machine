package bytemachine_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
)

func TestPrograms(t *testing.T) {
	t.Run("simple add", func(t *testing.T) {
		program := []byte{
			0x10, 0x02,
			0x10, 0x05,
			0x30,
			0xFF,
		}

		bm := bytemachine.NewByteMachine(program)
		bm.Run()

		require.Len(t, bm.GetStack(), 1, "stack should have exactly one value after addition")
		require.Equal(t, 7, bm.GetStack()[0], "2 + 5 should equal 7")
		require.True(t, bm.GetHalted(), "machine should be halted")
	})

	t.Run("multiple adds", func(t *testing.T) {
		program := []byte{
			0x10, 0x02,
			0x10, 0x08,
			0x30,
			0x10, 0x0F,
			0x30,
			0x10, 0x30,
			0x30,
			0xFF,
		}

		bm := bytemachine.NewByteMachine(program)
		bm.Run()

		require.Equal(t, 73, bm.GetStack()[0])
		require.True(t, bm.GetHalted(), "machine should be halted")
	})

	t.Run("store and load", func(t *testing.T) {
		program := []byte{
			0x10, 0x05,
			0x13, 0x01,
			0x10, 0x02,
			0x13, 0x02,
			0x14, 0x01,
			0xFF,
		}

		bm := bytemachine.NewByteMachine(program)
		bm.Run()

		require.Equal(t, 5, bm.GetStack()[0])
		require.Equal(t, 5, bm.GetRegisters()[1])
		require.Equal(t, 2, bm.GetRegisters()[2])
		require.True(t, bm.GetHalted(), "machine should be halted")
	})
}
