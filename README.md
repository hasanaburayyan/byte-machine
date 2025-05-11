# ByteMachine 🧠💾

ByteMachine is a small educational virtual machine built from scratch, designed to execute custom bytecode programs.

Inspired by projects like "Nand to Tetris," ByteMachine helps students and engineers understand how real CPUs, VMs, and low-level systems work — from memory management to instruction execution — but at a manageable, fun scale!

## ✨ Features

- Stack-based execution model
- 8 general-purpose registers
- Support for basic arithmetic (`ADD`, `SUB`, `MUL`, `DIV`, `MOD`, `POW`)
- Variable storage and retrieval (`STORE`, `LOAD`)
- Simple I/O (`OUT` instruction)
- Full custom opcode set with clean extensibility
- Instruction Pointer (IP) tracking and control flow
- Bytecode loaded from `.bin` files
- Assembly-to-binary compiler with label support
- Unit tests for each opcode at the `Apply()` level

## 🚀 How It Works

1. Write a `.bin` file containing raw bytecode instructions.
2. Load and execute the program using the `byte_machine`.
3. The virtual machine handles stack, registers, and control flow.
4. The program halts cleanly on the `HALT` instruction (`0xFF`).

### Example: Add two numbers (2 + 5)

| Byte | Meaning |
|:-----|:--------|
| 0x10 | `PUSH`  |
| 0x02 | 2       |
| 0x10 | `PUSH`  |
| 0x05 | 5       |
| 0x30 | `ADD`   |
| 0xFF | `HALT`  |

## 📦 Project Structure

```plaintext
src/
  cmd/byte_machine/       # CLI entry point (runs compiled .bin programs)
  cmd/debugger/           # Optional debugger (WIP)
  cmd/assembler/          # Assembler CLI for .bm source
  internal/
    byte_machine/         # Core VM: memory, stack, IP, execution
      machine/            # Machine abstraction
      opcodes/            # Opcode logic + decoding
        tests/            # Unit tests per opcode
    assembler/            # .bm to .bin conversion with label support
    utils/                # Helpers
program.bin                # Example compiled program
```

## 🛠️ Install

### ByteMachine

```bash
go install github.com/hasanaburayyan/byte-machine/src/cmd/byte_machine@latest
```

Best paired with [bytewrite](https://github.com/hasanaburayyan/bytewrite)!!

```bash
go install github.com/hasanaburayyan/bytewrite/src/cmd/bytewrite@latest
```

## 🖊️ Writing Programs in Assembly

You can write `.bm` files using readable opcodes and labels. Example:

### `print10.bm`

```asm
# Print numbers 1 through 10

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
```

## 🧪 Running Programs

### Option 1: Run a `.bm` file directly

```bash
go run cmd/assembler/main.go print10.bm --run
```

### Option 2: Compile `.bm` to `.bin`

```bash
go run cmd/assembler/main.go print10.bm
# Creates print10.bin
```

### Option 3: Run a `.bin` file

```bash
go run cmd/byte_machine/main.go print10.bin
```

## ⚙️ Byte-Level Examples (Raw Byte Execution)

If you want to skip assembly and just run raw bytes, use `bytewrite`:

### Print numbers 1–10

```bash
bytewrite -b   00010000 00000001   00010011 00000001   00010100 00000001   00010000 00001010   00100100   00010111 00010110   00010100 00000001   00000001   00010000 00000001   00110000   00010011 00000001   00010101 00000100   11111111 | byte_machine
```

### Print numbers 1–100

```bash
bytewrite -b   00010000 00000001   00010011 00000001   00010100 00000001   00010000 01100100   00100100   00010111 00010110   00010100 00000001   00000001   00010000 00000001   00110000   00010011 00000001   00010101 00000100   11111111 | byte_machine
```

## 🔜 Coming Soon

- Visual Studio Code debugger (DAP-based)
- Breakpoint support
- Decompiler: `.bin` → `.bm`
- Debugger-friendly pseudo-instructions

## 🤝 Contributing

PRs welcome! Feel free to open issues, contribute docs, or submit test programs. ByteMachine is a learning sandbox — let's make it fun and useful together.

## 📚 License

MIT © 2025 Hasan Abu-Rayyan
