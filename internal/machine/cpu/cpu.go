package cpu

import (
	"github.com/shotowon/giskv/internal/machine/bus"
)

type Register uint8

const (
	R_zero Register = iota // x0 ZERO register
	R_ra                   // x1 Return address of function calls
	R_sp                   // x2 Stack pointer register
	R_gp                   // x3 global pointer to global data
	R_tp                   // x4 thread pointer for thread local data
	// tmp registers t0-t6 are x5-7,x28-31
	// temporary register for short-lived values; not saved accross function calls
	R_t0 // x5
	R_t1 // x6
	R_t2 // x7
	// saved registers preserved accross function calls by the callee. x8-x9,x18-x27
	R_s0 // x8 Saved x Frame pointer
	R_s1 // x9
	// function arguments and return values (a0-7) x10-x17
	R_a0  // x10
	R_a1  // x11
	R_a2  // x12
	R_a3  // x13
	R_a4  // x14
	R_a5  // x15
	R_a6  // x16
	R_a7  // x17
	R_s2  // x18
	R_s3  // x19
	R_s4  // x20
	R_s5  // x21
	R_s6  // x22
	R_s7  // x23
	R_s8  // x24
	R_s9  // x25
	R_s10 // x26
	R_s11 // x27
	R_t3  // x28
	R_t4  // x29
	R_t5  // x30
	R_t6  // x31
)

type CPU struct {
	X   [32]uint32
	PC  uint32
	bus *bus.Bus
}

func New(bus *bus.Bus) *CPU {
	cpu := new(CPU)
	cpu.bus = bus
	cpu.X[R_zero] = 0
	cpu.PC = 0
	return cpu
}

func (c *CPU) Cycle() {
	instruction := c.Fetch32()
	opcode := instruction & 0b1111111
	fn3 := (instruction >> 12) & 0b111
	switch opcode {
	// i-type
	case 0b0010011:
		switch fn3 {
		case 0b000:
			c.addi(instruction)
		}
	case 0b0110011:
		fn7 := instruction >> 25
		switch fn3 {
		case 0b000:
			switch fn7 {
			case 0b00000000:
				c.add(instruction)
			case 0b01000000:
				c.sub(instruction)
			}
		case 0b001:
			switch fn7 {
			case 0b00000000:
				c.sll(instruction)
			}
		}
	}
	c.X[0] = uint32(0)
	c.PC += 4
}

func (c *CPU) Fetch32() uint32 {
	instruction := c.bus.Read32(c.PC)
	return instruction
}

// ADDI
// ADD-I instruction, rd = value(rs1) + signed(imm)
// rd - dst
// rs1 - src
// imm - immediate value in instruction
// [imm[11:0] | rs1 |fn3| rd  | opcode]
//	12         5    3    5      7

func (c *CPU) addi(instruction uint32) {
	rd := (instruction >> 7) & 0b11111
	rs1 := (instruction >> 15) & 0b11111
	imm := int32(instruction) >> 20
	src := int32(c.X[rs1]) + imm
	c.X[rd] = uint32(src)
}

func (c *CPU) add(instruction uint32) {
	rd := (instruction >> 7) & 0b11111
	rs1 := (instruction >> 15) & 0b11111
	rs2 := (instruction >> 20) & 0b11111
	c.X[rd] = c.X[rs1] + c.X[rs2]
}

func (c *CPU) sub(instruction uint32) {
	rd := (instruction >> 7) & 0b11111
	rs1 := (instruction >> 15) & 0b11111
	rs2 := (instruction >> 20) & 0b11111
	c.X[rd] = c.X[rs1] - c.X[rs2]
}

func (c *CPU) sll(instruction uint32) {
	rd := (instruction >> 7) & 0b11111
	rs1 := (instruction >> 15) & 0b11111
	rs2 := (instruction >> 20) & 0b11111
	c.X[rd] = c.X[rs1] << (c.X[rs2] & 0b11111)
}
