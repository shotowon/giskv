AS = riscv64-elf-as
OBJCOPY = riscv64-elf-objcopy

TEST_DIR = internal/machine/cpu/tests

build-instruction-test:
	$(AS) -march=rv32i -mabi=ilp32 $(TEST_DIR)/$(word 2,$(MAKECMDGOALS)).s -o $(TEST_DIR)/$(word 2,$(MAKECMDGOALS)).o
	$(OBJCOPY) -O binary $(TEST_DIR)/$(word 2,$(MAKECMDGOALS)).o $(TEST_DIR)/$(word 2,$(MAKECMDGOALS)).bin

%:
	@:
