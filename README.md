# ByteMachine 🧠💾

ByteMachine is a small educational virtual machine built from scratch, designed to execute custom bytecode programs.

Inspired by projects like "Nand to Tetris," ByteMachine helps students and engineers understand how real CPUs, VMs, and low-level systems work — from memory management to instruction execution — but at a manageable, fun scale!

---

## ✨ Features

- Stack-based execution model
- 8 general-purpose registers
- Support for basic arithmetic (ADD, SUB, MUL, DIV, MOD, POW)
- Variable storage and retrieval (STORE / LOAD)
- Simple I/O (OUT instruction)
- Full custom opcode set with clean extensibility
- Instruction Pointer (IP) tracking and control flow
- Bytecode loaded from `.bin` files
- Tests for each opcode (unit tested at Apply() level)

---

## 🚀 How It Works

1. Write a small binary file (`program.bin`) with your bytecode instructions.
2. Use the `ByteMachine` to load and execute the bytecode.
3. Stack, Registers, and Memory operate according to defined opcodes.
4. Program halts cleanly with `HALT` instruction (`0xFF`).

Example program that adds two numbers (2 + 5):

| Byte | Meaning |
|:---|:---|
| 0x10 | PUSH |
| 0x02 | 2 |
| 0x10 | PUSH |
| 0x05 | 5 |
| 0x30 | ADD |
| 0xFF | HALT |

---

## 📦 Project Structure

```plaintext
src/
  cmd/byte_machine/       # CLI entry point (loads and runs a .bin program)
  internal/
    byte_machine/         # Core VM (memory, IP, stack, registers)
      machine/            # Machine interface abstraction
      opcodes/            # Opcode definitions and Apply() logic
        tests/            # Unit tests for each opcode
    utils/                # Utility helpers (parsing, etc.)
program.bin                # Example compiled program
