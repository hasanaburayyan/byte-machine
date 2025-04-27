package bytemachine_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	bytemachine "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine"
)

func TestPrograms(t *testing.T) {
	t.Run("math program", func(t *testing.T) {
		// PUSH 2
		// PUSH 5
		// ADD         ; 2 + 5 = 7
		// PUSH 3
		// PUSH 1
		// SUB         ; 3 - 1 = 2
		// MUL         ; 7 * 2 = 14
		// HALT
		program := []byte{
			0x10, 0x02,
			0x10, 0x05,
			0x30,
			0x10, 0x03,
			0x10, 0x01,
			0x31,
			0x32,
			0xFF,
		}

		bm := bytemachine.NewByteMachine(program)
		bm.Run()

		require.True(t, bm.GetHalted())
		require.Equal(t, []int{14}, bm.GetStack())
	})

	t.Run("simple branching", func(t *testing.T) {
		// PUSH 0
		// JUMP_IF_ZERO 06
		// PUSH 99       ; (should skip this if jump happens)
		// PUSH 42       ; correct jump target
		// HALT

		program := []byte{
			0x10, 0x00,
			0x16, 0x06,
			0x10, 0x63,
			0x10, 0x2A,
			0xFF,
		}

		bm := bytemachine.NewByteMachine(program)
		bm.Run()

		require.True(t, bm.GetHalted())
		require.Equal(t, []int{42}, bm.GetStack())
	})

	t.Run("compare and branch", func(t *testing.T) {
		// PUSH 7
		// PUSH 5
		// GREATER      ; (7 > 5) → 1 pushed
		// JUMP_IF_NOT_ZERO 09
		// PUSH 99      ; should be skipped if greater
		// PUSH 42      ; correct jump target
		// HALT
		program := []byte{
			0x10, 0x07,
			0x10, 0x05,
			0x24,
			0x17, 0x09,
			0x10, 0x63,
			0x10, 0x2A,
			0xFF,
		}

		bm := bytemachine.NewByteMachine(program)
		bm.Run()

		require.True(t, bm.GetHalted())
		require.Equal(t, []int{42}, bm.GetStack())
	})

	t.Run("store and load math", func(t *testing.T) {
		// PUSH 6
		// STORE R0
		// PUSH 4
		// STORE R1
		// LOAD R0
		// LOAD R1
		// MUL
		// HALT

		program := []byte{
			0x10, 0x06,
			0x13, 0x00,
			0x10, 0x04,
			0x13, 0x01,
			0x14, 0x00,
			0x14, 0x01,
			0x32,
			0xFF,
		}

		bm := bytemachine.NewByteMachine(program)
		bm.Run()

		require.True(t, bm.GetHalted())
		require.Equal(t, []int{24}, bm.GetStack())
	})

	t.Run("loop test", func(t *testing.T) {
		// counter = 3
		// while counter != 0:
		// 		counter = counter - 1
		program := []byte{
			0x10, // (0) PUSH
			0x03, // (1) 3
			0x13, // (2) STORE
			0x00, // (3) R0
			0x14, // (4) LOAD
			0x00, // (5) R0
			0x10, // (6) PUSH
			0x00, // (7) 0
			0x20, // (8) EQ
			0x17, // (9) JUMP_IF_NOT_ZERO
			0x14, // (10) to HALT (address 20)
			0x14, // (11) LOAD
			0x00, // (12) R0
			0x10, // (13) PUSH
			0x01, // (14) 1
			0x31, // (15) SUB
			0x13, // (16) STORE
			0x00, // (17) R0
			0x15, // (18) JUMP
			0x04, // (19) to LOAD ()
			0xFF, // (20) HALT
		}

		bm := bytemachine.NewByteMachine(program)
		bm.Run()

		require.True(t, bm.GetHalted())
		require.Equal(t, 0, bm.GetRegisters()[0], "Counter should reach 0")
	})

	t.Run("nested looping", func(t *testing.T) {
		// outer_counter = 2
		// while outer_counter != 0:
		// 		inner_counter = 3
		// 		while inner_counter != 0:
		// 				inner_counter = inner_counter - 1
		// 		outer_counter = outer_counter - 1
		program := []byte{
			// Outer loop setup
			0x10, 0x02, // (0, 1) PUSH 2
			0x13, 0x00, // (2, 3) STORE R0 (outer counter)

			// -- Outer loop start --
			0x14, 0x00, // (4, 5) LOAD R0
			0x10, 0x00, // (6, 7) PUSH 0
			0x20,       // (8) EQ
			0x17, 0x28, // (9, 10) JUMP_IF_NOT_ZERO to HALT

			// Inner loop setup
			0x10, 0x03, // (11, 12) PUSH 3
			0x13, 0x01, // (13, 14) STORE R1 (inner counter)

			// -- Inner loop start --
			0x14, 0x01, // (15, 16) LOAD R1
			0x10, 0x00, // (17, 18) PUSH 0
			0x20,       // (19) EQ
			0x17, 0x1F, // (20, 21) JUMP_IF_NOT_ZERO to inner done

			// Inner body
			0x14, 0x01, // (22, 23) LOAD R1
			0x10, 0x01, // (24, 25) PUSH 1
			0x31,       // (26) SUB
			0x13, 0x01, // (27, 28) STORE R1
			0x15, 0x0F, // (29, 30) JUMP back to inner loop start (address 12)

			// -- Inner done --
			0x14, 0x00, // (31, 32) LOAD R0
			0x10, 0x01, // (33, 34) PUSH 1
			0x31,       // (35) SUB
			0x13, 0x00, // (36, 37) STORE R0
			0x15, 0x04, // (38, 39) JUMP back to outer loop start (address 4)

			// -- HALT (address 40) --
			0xFF, // (40)
		}

		bm := bytemachine.NewByteMachine(program)
		bm.Run()

		require.True(t, bm.GetHalted())
		require.Equal(t, 0, bm.GetRegisters()[0], "Outer counter should reach 0")
		require.Equal(t, 0, bm.GetRegisters()[1], "Inner counter should reach 0")
	})

	t.Run("loop with inc", func(t *testing.T) {
		program := []byte{
			0x10, 0x01,
			0x13, 0x01,
			0x14, 0x01,
			0x10, 0x0A,
			0x24, 0x17,
			0x16, 0x14,
			0x01, 0x01,
			0x10, 0x01,
			0x30, 0x13,
			0x01, 0x15,
			0x04, 0xFF,
		}
		bm := bytemachine.NewByteMachine(program)
		bm.Run()
	})
}
