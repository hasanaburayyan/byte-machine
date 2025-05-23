package assembler

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssembleValidProgram(t *testing.T) {

	t.Run("Simple add program", func(t *testing.T) {
		src := `PUSH 5
		PUSH 10
		ADD
		OUT
		HALT
	`

		// Assemble source
		output, sourceMap, err := Assemble(strings.NewReader(src))
		if err != nil {
			t.Fatalf("assemble failed: %v", err)
		}

		// Expected binary output (replace these with actual byte values from your opcode map)
		expected := []byte{
			0x10, 0x05, // PUSH 5
			0x10, 0x0A, // PUSH 10
			0x30, // ADD
			0x01, // OUT
			0xFF, // HALT
		}

		require.Equal(t, expected, output)
		require.Equal(t, 7, len(sourceMap.ByteToLine))
		require.Equal(t, 5, len(sourceMap.LineToByte)) // each instruction on its own line
		require.Equal(t, 1, sourceMap.ByteToLine[0])   // PUSH 5 opcode
		require.Equal(t, 1, sourceMap.ByteToLine[1])   // PUSH 5 arg
		require.Equal(t, 2, sourceMap.ByteToLine[2])   // PUSH 10 opcode
		require.Equal(t, 2, sourceMap.ByteToLine[3])   // PUSH 10 arg
		require.Equal(t, 3, sourceMap.ByteToLine[4])   // ADD
	})

	t.Run("bad opcode", func(t *testing.T) {
		src := `FROBNICATE`

		_, _, err := Assemble(strings.NewReader(src))

		require.Error(t, err)
	})

	t.Run("not enough arg count", func(t *testing.T) {
		src := `PUSH`

		_, _, err := Assemble(strings.NewReader(src))

		require.Error(t, err)
	})

	t.Run("too many enough arg count", func(t *testing.T) {
		src := `PUSH 10 11`

		_, _, err := Assemble(strings.NewReader(src))

		require.Error(t, err)
	})

	t.Run("looping", func(t *testing.T) {
		src := `
		PUSH 1
		STORE 1

		loop:
		LOAD 1
		PUSH 10
		GREATER
		JUMP_IF_NOT_ZERO end

		LOAD 1
		OUT
		PUSH 1
		ADD
		STORE 1
		JUMP loop

		end:
		HALT

		`
		expected := []byte{
			0x10, 0x01,
			0x13, 0x01,
			0x14, 0x01,
			0x10, 0x0a,
			0x24, 0x17,
			0x15, 0x14,
			0x01, 0x01,
			0x10, 0x01,
			0x30, 0x13,
			0x01, 0x15,
			0x04, 0xff,
		}

		output, sourceMap, err := Assemble(strings.NewReader(src))

		require.NoError(t, err)
		require.Equal(t, expected, output)
		require.Equal(t, 22, len(sourceMap.ByteToLine))
		require.Equal(t, 13, len(sourceMap.LineToByte))
	})
}
