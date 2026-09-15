package bus

import "fmt"

type Bus struct {
	RAM []byte
}

func Load(program []byte) *Bus {
	ram := make([]byte, len(program)+4*1024)
	copy(ram, program)
	return &Bus{
		RAM: ram,
	}
}

func (b *Bus) Read8(addr uint32) uint8 {
	if int(addr) >= len(b.RAM) {
		panic(fmt.Sprintf("bus: invalid memory read: crashing, addr: uint32(%d) hex(%x)", addr, addr))
	}

	return b.RAM[addr]
}

func (b *Bus) Write8(addr uint32, value uint8) {
	if int(addr) >= len(b.RAM) {
		panic(fmt.Sprintf("invalid memory write: crashing, addr: uint32(%d) hex(%x), value: uint8(%d) hex(%x)", addr, addr, value, value))
	}

	b.RAM[addr] = value
}

func (b *Bus) Read32(addr uint32) uint32 {
	if int(addr)+4 > len(b.RAM) {
		panic(fmt.Sprintf("bus: invalid memory read: crashing, addr: uint32(%d) hex(%x)", addr, addr))
	}

	return uint32(b.RAM[addr]) | uint32(b.RAM[addr+1])<<8 | uint32(b.RAM[addr+2])<<16 | uint32(b.RAM[addr+3])<<24
}

func (b *Bus) Write32(addr uint32, value uint32) {
	if int(addr)+4 > len(b.RAM) {
		panic(fmt.Sprintf("bus: invalid memory write: crashing, addr: uint32(%d) hex(%x)", addr, addr))
	}

	b.RAM[addr] = byte(value)
	b.RAM[addr+1] = byte(value >> 8)
	b.RAM[addr+2] = byte(value >> 16)
	b.RAM[addr+3] = byte(value >> 24)
}
