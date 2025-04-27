package utils

import "encoding/binary"

func ParseInt(parts []byte) int {
	switch len(parts) {
	case 1:
		return int(parts[0])
	case 2:
		return int(binary.BigEndian.Uint16(parts))
	case 4:
		return int(binary.BigEndian.Uint32(parts))
	default:
		panic("unsupported number of bytes to ParseInt")
	}
}
