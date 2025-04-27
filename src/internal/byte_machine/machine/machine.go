package machine

type Machine interface {
	GetMemory() []byte
	GetIP() int
	IncrementIP()
	DecrementIP()
	SetIP(ip int)
	GetStack() []int
	Push(val int)
	Peek() int
	Pop() int
	GetRegisters() [8]int
	SetRegister(reg int, val int)
	GetHalted() bool
	Halt()
}
