package machine

import (
	"github.com/shotowon/giskv/internal/machine/bus"
	"github.com/shotowon/giskv/internal/machine/cpu"
)

type Machine struct {
	CPU *cpu.CPU
	bus *bus.Bus
}

func New(cpu *cpu.CPU, bus *bus.Bus) *Machine {
	return &Machine{
		CPU: cpu,
		bus: bus,
	}
}

func (m *Machine) Run() error {
	for int(m.CPU.PC+4) < len(m.bus.RAM) {
		m.CPU.Cycle()
	}
	return nil
}
