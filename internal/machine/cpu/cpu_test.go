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
