package opcodes

import "github.com/hasanaburayyan/byte-machine/src/internal/byte_machine/machine"

type OpCode interface {
	Apply(m machine.Machine)
}

type OpCodeDef struct {
	Name        string
	Description string
	Op          OpCode
}

var OpCodeDefs = map[byte]OpCodeDef{
	0x00: {Name: "NO_OP", Description: "Does nothing", Op: &NOOP{}},
	0x01: {Name: "OUT", Description: "Prints outs top of stack", Op: &Out{}},
	// 0x02 - 0x0F are reserved for future
	0x10: {Name: "PUSH", Description: "Puts value onto stack", Op: &Push{}},
	0x11: {Name: "POP", Description: "Pops value from stack", Op: &Pop{}},
	0x12: {Name: "PEEK", Description: "Puts top of stack on register without removing", Op: &UnImplemented{}},
	0x13: {Name: "STORE", Description: "Stores value to register", Op: &Store{}},
	0x14: {Name: "LOAD", Description: "Loads value from register", Op: &Load{}},
	0x15: {Name: "JUMP", Description: "Jumps to instruction", Op: &Jump{}},
	0x16: {Name: "JUMP_IF_ZERO", Description: "Jump if false", Op: &JumpIfZero{}},
	0x17: {Name: "JUMP_IF_NOT_ZERO", Description: "Jump if true", Op: &JumpIfNotZero{}},
	// 0x18 - 0x1F are reserved for more control flow
	0x20: {Name: "EQ", Description: "Checks equality and puts 0 on register if false, 1 if true", Op: &Eq{}},
	0x21: {Name: "NE", Description: "Checks NOT equal to and puts 0 on register if false, 1 if true", Op: &Ne{}},
	0x22: {Name: "LESS", Description: "Checks less than and puts 0 on register if false, 1 if true", Op: &Less{}},
	0x23: {Name: "LESS_OR_EQ", Description: "Checks less than or equal to and puts 0 on register if false, 1 if true", Op: &LessOrEq{}},
	0x24: {Name: "GREATER", Description: "Checks greater than and puts 0 on register if false, 1 if true", Op: &Greater{}},
	0x25: {Name: "GREATER_OR_EQ", Description: "Checks greater than or equal and puts 0 on register if false, 1 if true", Op: &GreaterOrEq{}},
	// 0x26 - 0x2F are reserved for more compare functions
	0x30: {Name: "ADD", Description: "Adds two values from stack and places result back on stack", Op: &Add{}},
	0x31: {Name: "SUB", Description: "Subtracts two values from stack and places result back on stack (first pop is right operand)", Op: &Subtract{}},
	0x32: {Name: "MUL", Description: "Multiplies two values from stack and places result back on stack", Op: &Multiply{}},
	0x33: {Name: "DIV", Description: "Divides two values from stack and places result back on stack (first pop is right operand)", Op: &Divide{}},
	0x34: {Name: "MOD", Description: "Modulo two vales from stack and places result back on stack (first pop is right operand)", Op: &Modulo{}},
	0x35: {Name: "POW", Description: "Raises a value to a power (first pop is exponent)", Op: &Power{}},
	0xFF: {Name: "HALT", Description: "Stops program", Op: &Halt{}},
}
