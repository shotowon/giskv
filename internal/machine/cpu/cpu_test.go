package cpu_test

import (
	"os"
	"testing"

	"github.com/shotowon/giskv/internal/machine/bus"
	"github.com/shotowon/giskv/internal/machine/cpu"
)

func TestInstructions(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		cycles   int
		reg      cpu.Register
		expected uint32
	}{
		{
			name:     "ADDI",
			file:     "tests/addi.bin",
			cycles:   1,
			reg:      cpu.R_t0,
			expected: 42,
		},
		{
			name:     "SLTI",
			file:     "tests/slti.bin",
			cycles:   2,
			reg:      cpu.R_t2,
			expected: 1,
		},
		{
			name:     "SLTIU",
			file:     "tests/sltiu.bin",
			cycles:   2,
			reg:      cpu.R_t2,
			expected: 0,
		},
		{
			name:     "XORI",
			file:     "tests/xori.bin",
			cycles:   2,
			reg:      cpu.R_t2,
			expected: 0b1011,
		},
		{
			name:     "ORI",
			file:     "tests/ori.bin",
			cycles:   2,
			reg:      cpu.R_t2,
			expected: 0b1111,
		},
		{
			name:     "ANDI",
			file:     "tests/andi.bin",
			cycles:   2,
			reg:      cpu.R_t2,
			expected: 0b1000001,
		},
		{
			name:     "SLLI",
			file:     "tests/slli.bin",
			cycles:   2,
			reg:      cpu.R_t2,
			expected: 0b11010010,
		},
		{
			name:     "SRLI",
			file:     "tests/srli.bin",
			cycles:   2,
			reg:      cpu.R_t2,
			expected: 0b110100,
		},
		{
			name:     "SRAI",
			file:     "tests/srai.bin",
			cycles:   2,
			reg:      cpu.R_t2,
			expected: 0b11111111111111111111111111110000,
		},
		{
			name:     "ADD",
			file:     "tests/add.bin",
			cycles:   3,
			reg:      cpu.R_t2,
			expected: 55,
		},
		{
			name:     "SUB",
			file:     "tests/sub.bin",
			cycles:   3,
			reg:      cpu.R_t2,
			expected: 29,
		},
		{
			name:     "SLL",
			file:     "tests/sll.bin",
			cycles:   3,
			reg:      cpu.R_t2,
			expected: 12,
		},
		{
			name:     "SRL",
			file:     "tests/srl.bin",
			cycles:   3,
			reg:      cpu.R_t2,
			expected: 1,
		},
		{
			name:     "SRA",
			file:     "tests/sra.bin",
			cycles:   3,
			reg:      cpu.R_t2,
			expected: 0b11111111111111111111111111111110,
		},
		{
			name:     "SLT",
			file:     "tests/slt.bin",
			cycles:   3,
			reg:      cpu.R_t2,
			expected: 1,
		},
		{
			name:     "SLTU",
			file:     "tests/sltu.bin",
			cycles:   3,
			reg:      cpu.R_t2,
			expected: 0,
		},
		{
			name:     "XOR",
			file:     "tests/xor.bin",
			cycles:   3,
			reg:      cpu.R_t2,
			expected: 4,
		},
		{
			name:     "OR",
			file:     "tests/or.bin",
			cycles:   3,
			reg:      cpu.R_t2,
			expected: 0b1111,
		},
		{
			name:     "AND",
			file:     "tests/and.bin",
			cycles:   3,
			reg:      cpu.R_t2,
			expected: 0b1001,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.ReadFile(tt.file)
			if err != nil {
				t.Fatalf("failed to read test: %v", err)
			}

			b := bus.Load(f)
			c := cpu.New(b)

			for range tt.cycles {
				c.Cycle()
			}

			got := c.X[tt.reg]

			if got != tt.expected {
				t.Fatalf(
					"expected x%d to be: %d, got: %d",
					tt.reg,
					tt.expected,
					got,
				)
			}
		})
	}
}

func TestLB(t *testing.T) {
	f, err := os.ReadFile("tests/lb.bin")
	if err != nil {
		t.Fatalf("failed to read test: %v", err)
	}

	b := bus.Load(f)
	c := cpu.New(b)

	c.Cycle()
	c.Cycle()

	if c.X[7] != 0xffffffe4 {
		t.Fatalf(
			"expected x7 == 0xffffffe4, got = %x",
			c.X[7],
		)
	}

	if uint8(c.X[7]) != 0xe4 {
		t.Fatalf(
			"expected x7 == 0xe4, got = %x",
			uint8(c.X[7]),
		)
	}
}

func TestLBU(t *testing.T) {
	f, err := os.ReadFile("tests/lbu.bin")
	if err != nil {
		t.Fatalf("failed to read test: %v", err)
	}

	b := bus.Load(f)
	c := cpu.New(b)

	c.Cycle()
	c.Cycle()

	if c.X[7] != 0x000000e4 {
		t.Fatalf(
			"expected x7 == 0x000000e4, got = %x",
			c.X[7],
		)
	}

	if uint8(c.X[7]) != 0xe4 {
		t.Fatalf(
			"expected x7 == 0xe4, got = %x",
			uint8(c.X[7]),
		)
	}
}

func TestLH(t *testing.T) {
	f, err := os.ReadFile("tests/lh.bin")
	if err != nil {
		t.Fatalf("failed to read test: %v", err)
	}

	b := bus.Load(f)
	c := cpu.New(b)

	c.Cycle()
	c.Cycle()

	if c.X[7] != 0xfffff43f {
		t.Fatalf(
			"expected x7 == fffff43f, got = %x",
			uint16(c.X[7]),
		)
	}

	if uint16(c.X[7]) != 0xf43f {
		t.Fatalf(
			"expected x7 == f43f, got = %x",
			uint16(c.X[7]),
		)
	}
}

func TestLW(t *testing.T) {
	f, err := os.ReadFile("tests/lw.bin")
	if err != nil {
		t.Fatalf("failed to read test: %v", err)
	}

	b := bus.Load(f)
	c := cpu.New(b)

	c.Cycle()
	c.Cycle()

	if c.X[7] != 0xa5c3f43f {
		t.Fatalf(
			"expected x7 == a5c3f43f, got = %x",
			c.X[7],
		)
	}
}
