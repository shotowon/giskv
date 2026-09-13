package cpu_test

import (
	"os"
	"testing"

	"github.com/shotowon/giskv/internal/machine/bus"
	"github.com/shotowon/giskv/internal/machine/cpu"
)

func TestInstructions(t *testing.T) {
}

func TestADDI(t *testing.T) {
	f, err := os.ReadFile("tests/addi.bin")
	if err != nil {
		t.Fatalf("failed to read test for add with error: %v", err)
	}
	b := bus.Load(f)
	c := cpu.New(b)
	c.Cycle()

	expected := uint32(42)
	got := c.X[cpu.R_t0]

	if got != expected {
		t.Fatalf("expected x5 to be: %d, got: %d", expected, got)
	}
}

func ADD(t *testing.T) {
	f, err := os.ReadFile("tests/add.bin")
	if err != nil {
		t.Fatalf("failed to read test for add with error: %v", err)
	}
	b := bus.Load(f)
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

func TestSLL(t *testing.T) {
	f, err := os.ReadFile("tests/sll.bin")
	if err != nil {
		t.Fatalf("failed to read test for add with error: %v", err)
	}
	b := bus.Load(f)

	c := cpu.New(b)

	c.Cycle()
	c.Cycle()
	c.Cycle()

	expected := uint32(12)
	got := c.X[cpu.R_t2]

	if got != expected {
		t.Fatalf("expected x7 to be: %d, got: %d", expected, got)
	}
}
