package assembler

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/opcodes"
)

type op struct {
	BValue   byte
	ArgCount int
}

type SourceMap struct {
	ByteToLine map[int]int
	LineToByte map[int][]int
}

func createOpCodeMap() map[string]op {
	m := make(map[string]op)
	for key, val := range opcodes.OpCodeDefs {
		m[val.Name] = op{
			BValue:   key,
			ArgCount: val.ArgCount,
		}
	}
	return m
}

// First pass: create label map
func createLabelMap(lines []string, ops map[string]op) (map[string]int, error) {
	labels := make(map[string]int)
	byteOffset := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		tokens := strings.Fields(line)
		if len(tokens) == 0 {
			continue
		}

		first := tokens[0]
		if strings.HasSuffix(first, ":") {
			label := strings.TrimSuffix(first, ":")
			labels[label] = byteOffset
			continue
		}

		opcode, ok := ops[first]
		if !ok {
			return nil, fmt.Errorf("unknown opcode '%s' while building label map", first)
		}

		byteOffset++                  // for opcode
		byteOffset += opcode.ArgCount // for args
	}

	return labels, nil
}

// Second pass: generate bytecode
func Assemble(reader io.Reader) ([]byte, *SourceMap, error) {
	ops := createOpCodeMap()
	sourceMap := &SourceMap{
		ByteToLine: make(map[int]int),
		LineToByte: make(map[int][]int),
	}

	// Read all lines once
	scanner := bufio.NewScanner(reader)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	// First pass: map labels to byte offsets
	labels, err := createLabelMap(lines, ops)
	if err != nil {
		return nil, nil, err
	}

	// Second pass: assemble
	var output []byte
	byteOffset := 0

	for lineNum, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		tokens := strings.Fields(line)
		if len(tokens) == 0 {
			continue
		}

		opcodeName := tokens[0]

		// skip label definitions
		if strings.HasSuffix(opcodeName, ":") {
			continue
		}

		opcode, ok := ops[opcodeName]
		if !ok {
			return nil, nil, fmt.Errorf("unknown opcode '%s' on line %d", opcodeName, lineNum+1)
		}

		// Record mapping from current byteOffset BEFORE adding the opcode
		sourceMap.ByteToLine[byteOffset] = lineNum + 1
		sourceMap.LineToByte[lineNum+1] = append(sourceMap.LineToByte[lineNum+1], byteOffset)

		output = append(output, opcode.BValue)
		byteOffset++

		if opcode.ArgCount != len(tokens)-1 {
			return nil, nil, fmt.Errorf("expected %d args but got %d on line: %d", opcode.ArgCount, len(tokens)-1, lineNum+1)
		}

		for i := 0; i < opcode.ArgCount; i++ {
			arg := tokens[i+1]

			// try parsing as number
			val, err := strconv.Atoi(arg)
			if err == nil {
				output = append(output, byte(val))
				sourceMap.ByteToLine[byteOffset] = lineNum + 1
				sourceMap.LineToByte[lineNum+1] = append(sourceMap.LineToByte[lineNum+1], byteOffset)
				byteOffset++
				continue
			}

			// fallback to label
			addr, found := labels[arg]
			if !found {
				return nil, nil, fmt.Errorf("unknown label or invalid argument '%s' on line %d", arg, lineNum+1)
			}
			output = append(output, byte(addr))
			sourceMap.ByteToLine[byteOffset] = lineNum + 1
			sourceMap.LineToByte[lineNum+1] = append(sourceMap.LineToByte[lineNum+1], byteOffset)
			byteOffset++
		}
	}

	return output, sourceMap, nil
}
