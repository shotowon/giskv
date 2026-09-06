package cpu_test

import (
	"testing"

	"github.com/shotowon/giskv/internal/machine/bus"
	"github.com/shotowon/giskv/internal/machine/cpu"
)

func TestADDI(t *testing.T) {
	b := bus.Load([]byte{
		0b10010011, // opcode I-type
		0b00000010, // rd set to x5
		0b10100000, //
		0b00000010, // imm = 42
	})
	c := cpu.New(b)
	c.Cycle()

	expected := uint32(42)
	got := c.X[cpu.R_t0]

	if got != expected {
		t.Fatalf("expected x5 to be: %d, got: %d", expected, got)
	}
}

func TestADD(t *testing.T) {
	b := bus.Load([]byte{
		0b10010011,
		0b00000010,
		0b10100000,
		0b00000010, // x5 = 42 from addi test

		0b10010011,
		0b00000001,
		0b11010000,
		0b00000000, // x3 = 13 from addi test

		// x7 = x3 + x5
		0b10110011, // opcode R-Type
		0b10000011, // register 7 dst, funct3 000 ADD-SUB
		0b01010001, // add x3 + x5
		0b00000000,
	})
	c := cpu.New(b)
	c.Cycle()
	c.Cycle()
	c.Cycle()

	expected := uint32(55)
	got := c.X[cpu.R_t2]

	if got != expected {
		t.Fatalf("expected x7 to be: %d, got: %d", expected, got)
	}
}
